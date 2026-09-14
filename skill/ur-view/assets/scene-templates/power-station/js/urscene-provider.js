/* 文件说明：真实电表数据桥；仅访问宿主绑定设备，不保存凭证或生成演示数据。
 * 流程：校验绑定 → 按设备组合标识共享状态 → 限并发快照 → 属性推送 → 选中指标按需历史。
 * 绑定版本隔离在途请求；字段各自校验时间和有效期，历史缺失保留断点。
 */
(function () {
  'use strict';
  /** 固定展示字段及标准单位；配置负责把物模型原始单位换算为这些单位。 */
  const METRICS = [
    { field: 'p', label: '有功', unit: 'kW', decimals: 1 },
    { field: 'ia', label: '电流', unit: 'A', decimals: 1 },
    { field: 'ua', label: '电压', unit: 'V', decimals: 1 },
    { field: 'q', label: '无功', unit: 'kvar', decimals: 2 },
    { field: 'pf', label: '功率因数', unit: '', decimals: 3 },
    { field: 'kwh', label: '电能', unit: 'kWh', decimals: 1 }
  ];
  /** 严格解析数值，禁止空字符串、布尔值和无穷值变成有效读数。 */
  function number(value) {
    if (value == null || typeof value === 'boolean' || (typeof value === 'string' && !value.trim())) return null;
    const result = Number(value); // 合法字符串允许转换。
    return Number.isFinite(result) ? result : null;
  }
  /** 接受秒或毫秒时间戳；无时间或明显未来时间不能作为当前有效记录。 */
  function timestamp(value) {
    let result = number(value); // 标准化到毫秒。
    if (result === null || result <= 0) return 0;
    if (result < 1e12) result *= 1000;
    return result <= Date.now() + 60000 ? result : 0;
  }
  /** 设备组合身份避免不同产品同名设备互相覆盖。 */
  function key(device) { return JSON.stringify([String(device.productID), device.deviceName]); }
  /** 同一设备共享原始字段与历史缓存，锚点只提供视图。 */
  function empty() { return { fields: {}, history: {}, trendStatus: {}, historyAt: {}, error: false }; }
  /** 校验统一响应；失败不伪装为成功空记录。 */
  function dataOf(response) {
    if (!response || (response.code !== undefined && response.code !== 200)) throw new Error('数据查询失败');
    return response.code === 200 ? (response.data || {}) : response;
  }
  /** 完整响应与 SDK 已解包结果统一读取列表。 */
  function listOf(response) {
    const data = dataOf(response); // 统一业务负载。
    return Array.isArray(data.list) ? data.list : [];
  }
  /** 输入布局清单和可选接入配置，建立空状态。 */
  function UrSceneProvider(meters, options) {
    this.meters = meters;
    this.config = options || window.STATION_DATA;
    this.states = {};
    this.listeners = [];
    this.mode = 'pending';
    this.bindings = {};
    this.devices = {};
    this.records = {};
    this.generation = 0;
    this.pending = new Map();
    this.queue = [];
    this.active = 0;
    this.signature = null;
    this.sdk = window.UrSceneSDK || null;
    this.unsubscribers = [];
    this.stopped = false;
    this._notify();
  }
  /** 从物模型原值转换到约定展示单位；未映射字段不猜测别名。 */
  UrSceneProvider.prototype._convert = function (field, value) {
    const config = this.config.fields[field]; // 配置标识符和倍率。
    if (!config || !config.identifier) return null;
    if (config.type === 'boolean') {
      if (String(value) === String(config.trueValue)) return true;
      if (String(value) === String(config.falseValue)) return false;
      return null;
    }
    const raw = number(value), multiplier = number(config.multiplier); // 严格数值与单位换算倍率。
    if (raw === null || multiplier === null || multiplier <= 0) return null;
    return Number.isFinite(raw * multiplier) ? raw * multiplier : null;
  };
  /** 写入单字段；不允许旧消息或相同时间戳替换新记录。 */
  UrSceneProvider.prototype._apply = function (deviceKey, identifier, value, time) {
    const record = this.records[deviceKey], ts = timestamp(time); // 记录与标准时间。
    if (!record || !ts) return false;
    let changed = false; // 是否有可见字段更新。
    Object.entries(this.config.fields).forEach(([field, config]) => {
      if (!config.identifier || config.identifier !== identifier) return;
      const converted = this._convert(field, value), old = record.fields[field]; // 单位归一后的值及前值。
      if (converted === null || (old && old.ts >= ts)) return;
      record.fields[field] = { value: converted, ts };
      if (typeof converted === 'number') this._merge(record, field, [{ t: ts, v: converted }]);
      record.error = false;
      changed = true;
    });
    return changed;
  };
  /** 合并有效原始趋势点，实时已存在的同时间点优先。 */
  UrSceneProvider.prototype._merge = function (record, field, points) {
    const byTime = new Map(points.map(point => [point.t, point.v])); // 时间戳去重表。
    (record.history[field] || []).forEach(point => byTime.set(point.t, point.v));
    record.history[field] = Array.from(byTime, ([t, v]) => ({ t, v })).filter(point => point.t >= Date.now() - 86400000 - this.config.staleMs).sort((a, b) => a.t - b.t);
  };
  /** 以固定时间网格取当时之前的有效点，缺口返回 null，不补零或插值。 */
  UrSceneProvider.prototype._series = function (points, now) {
    if (!points || !points.length) return [];
    const step = this.config.historyStepMs, end = Math.floor(now / step) * step, result = []; // 时间网格范围。
    let cursor = -1; // 已走过的最新采样。
    for (let time = end - 86400000; time <= end; time += step) {
      while (cursor + 1 < points.length && points[cursor + 1].t <= time) cursor++;
      const point = points[cursor]; // 当时之前的一条记录。
      result.push({ t: time, v: point && time - point.t <= this.config.staleMs ? point.v : null });
    }
    return result;
  };
  /** 投影每个锚点的当前有效读数；过期字段单独置空，不由其他字段更新时间续命。 */
  UrSceneProvider.prototype._notify = function () {
    const now = Date.now(); // 此次统一有效期基准。
    this.meters.forEach(meter => {
      const record = this.records[this.bindings[meter.id]], state = { id: meter.id, ts: null, trends: {}, trendStatus: {} }; // 可见状态。
      let valid = 0, expired = 0; // 当前与过期字段计数。
      Object.keys(this.config.fields).forEach(field => {
        const value = record && record.fields[field]; // 字段独立时间戳。
        const fresh = value && now - value.ts <= this.config.staleMs; // 新鲜度判断。
        state[field] = fresh ? value.value : null;
        if (fresh) { valid++; state.ts = Math.max(state.ts || 0, value.ts); }
        else if (value) expired++;
      });
      state.status = !record ? 'unbound' : !Object.values(this.config.fields).some(config => config.identifier) ? 'unconfigured' : valid ? 'ready' : record.error ? 'error' : expired ? 'stale' : 'empty';
      if (record) {
        state.trendStatus = { ...record.trendStatus };
        Object.keys(record.history).forEach(field => { state.trends[field] = this._series(record.history[field], now); });
      }
      this.states[meter.id] = state;
    });
    this.listeners.forEach(listener => listener(this.states));
  };
  /** 所有网络查询共用三路队列；重复设备请求合并，在途旧版本禁止提交结果。 */
  UrSceneProvider.prototype._request = function (cacheKey, task) {
    const id = this.generation + ':' + cacheKey; // 绑定版本参与请求去重。
    if (this.pending.has(id)) return this.pending.get(id);
    const promise = new Promise(resolve => { this.queue.push({ id, generation: this.generation, task, resolve }); }); // 等待排队任务。
    this.pending.set(id, promise);
    this._drain();
    return promise;
  };
  /** 串行补充队列槽位；停用和旧绑定任务直接丢弃。 */
  UrSceneProvider.prototype._drain = function () {
    while (this.active < 3 && this.queue.length) {
      const job = this.queue.shift(); // 下一条任务。
      if (this.stopped || job.generation !== this.generation) { this.pending.delete(job.id); job.resolve(false); continue; }
      this.active++;
      Promise.resolve().then(job.task).catch(() => false).then(value => {
        this.active--;
        this.pending.delete(job.id);
        job.resolve(value);
        this._drain();
      });
    }
  };
  /** 重新绑定彻底清理旧映射与缓存；重复握手保留原有记录。 */
  UrSceneProvider.prototype._handleInit = function (bindings) {
    if (this.stopped) return;
    const known = new Set(this.meters.map(meter => meter.id)), normalized = []; // 只接受当前布局锚点。
    (bindings || []).forEach(binding => {
      const id = binding && typeof binding.path === 'string' && binding.path.startsWith('meter/') ? binding.path.slice(6) : ''; // 锚点本地编号。
      if (!known.has(id) || !binding.productID || typeof binding.deviceName !== 'string' || !binding.deviceName) return;
      normalized.push({ id, productID: binding.productID, deviceName: binding.deviceName });
    });
    normalized.sort((a, b) => JSON.stringify(a).localeCompare(JSON.stringify(b)));
    const signature = JSON.stringify(normalized); // 不依赖宿主绑定条目顺序。
    if (signature === this.signature) return;
    this.signature = signature;
    this.generation++;
    this.bindings = {}; this.devices = {}; this.records = {};
    const grouped = new Map(); // 同一锚点有多个不同设备时拒绝歧义绑定。
    normalized.forEach(binding => {
      const choices = grouped.get(binding.id) || new Map(); // 锚点对应设备集合。
      choices.set(key(binding), binding); grouped.set(binding.id, choices);
    });
    grouped.forEach((choices, id) => {
      if (choices.size !== 1) return;
      const device = Array.from(choices.values())[0], deviceKey = key(device); // 确定的组合身份。
      this.bindings[id] = deviceKey;
      this.devices[deviceKey] = device;
      this.records[deviceKey] = this.records[deviceKey] || empty();
    });
    this.mode = Object.keys(this.devices).length ? 'real' : 'unbound';
    if (this.onModeChange) this.onModeChange(this.mode);
    this._notify();
    this._refresh();
  };
  /** 只按已绑定唯一设备补快照；已核实汇总成员优先，历史保持按需。 */
  UrSceneProvider.prototype._refresh = function () {
    const ids = this.config.summary.verified ? this.config.summary.meterIds : [], seen = new Set(); // 优先顺序不隐含实际总分关系。
    [...ids, ...Object.keys(this.bindings)].forEach(id => {
      const deviceKey = this.bindings[id]; // 设备去重。
      if (!deviceKey || seen.has(deviceKey)) return;
      seen.add(deviceKey); this.fetchDetail(id);
    });
  };
  /** 消费宿主绑定初值或属性推送；产品身份缺失时拒绝裸设备名匹配。 */
  UrSceneProvider.prototype._handleData = function (payload) {
    if (this.stopped || !payload || this.mode !== 'real') return;
    const device = payload.device || payload, deviceKey = key(device); // 两种 SDK 消息统一入口。
    if (!this.records[deviceKey]) return;
    const params = payload.params || (payload.identifier ? { [payload.identifier]: payload.value !== undefined ? payload.value : payload.param } : {}); // 属性集合。
    let changed = false; // 一帧最多刷新一次 UI。
    Object.keys(params).forEach(identifier => { changed = this._apply(deviceKey, identifier, params[identifier], payload.ts || payload.timestamp) || changed; });
    if (changed) this._notify();
  };
  /** 启动只注册宿主 SDK，无宿主时保留未连接状态与完整模型。 */
  UrSceneProvider.prototype.start = function () {
    if (!this.sdk || window.parent === window) { this.mode = 'unbound'; if (this.onModeChange) this.onModeChange(this.mode); this._notify(); return; }
    const sdk = this.sdk; // 凭证始终由宿主处理。
    [sdk.onInit(bindings => this._handleInit(bindings)), sdk.onData(payload => this._handleData(payload)), sdk.subscribe('prop', payload => this._handleData(payload))].forEach(off => { if (typeof off === 'function') this.unsubscribers.push(off); });
    sdk.reportAnchors(this.meters.map(meter => ({ path: 'meter/' + meter.id, name: meter.name, desc: meter.cabinetName })));
    sdk.ready();
    this.timer = setInterval(() => { this._notify(); this._refresh(); }, this.config.pollMs);
  };
  /** 停止计时与后续请求；在途网络结果由版本号隔离。 */
  UrSceneProvider.prototype.stop = function () { this.stopped = true; this.generation++; clearInterval(this.timer); this.unsubscribers.forEach(off => off()); this._drain(); };
  /** 注册 UI 状态订阅，返回取消订阅函数。 */
  UrSceneProvider.prototype.subscribe = function (listener) { this.listeners.push(listener); return () => { this.listeners = this.listeners.filter(item => item !== listener); }; };
  /** 返回锚点当前状态。 */
  UrSceneProvider.prototype.get = function (id) { return this.states[id]; };
  /** 严格汇总经核验的进线成员；缺失任何成员返回空值，不把缺数据当零。 */
  UrSceneProvider.prototype.summary = function (field) {
    const relation = this.config.summary, unique = new Set(), registry = new Map(this.meters.map(meter => [meter.id, meter])); // 显式计量关系与设备去重。
    if (!relation.verified || !relation.meterIds.length) return null;
    let total = 0; // 有符号功率不截断负数。
    for (const id of relation.meterIds) {
      const meter = registry.get(id), deviceKey = this.bindings[id], state = this.states[id]; // 每个汇总成员必须绑定且属于进线。
      if (!meter || meter.kind !== 'incoming' || !deviceKey || !state || state[field] == null) return null;
      if (unique.has(deviceKey)) continue;
      unique.add(deviceKey); total += state[field];
    }
    return total;
  };
  /** 查询一台已绑定设备快照，字段独立处理时间戳、单位、有效期。 */
  UrSceneProvider.prototype.fetchDetail = function (id) {
    const deviceKey = this.bindings[id], device = this.devices[deviceKey], generation = this.generation; // 捕获绑定版本。
    if (!device || !this.sdk || !this.sdk.callApi || !Object.values(this.config.fields).some(config => config.identifier)) return Promise.resolve(false);
    return this._request(deviceKey + ':latest', async () => {
      try {
        const response = await this.sdk.callApi('/api/v1/things/device/msg/property-latest/get-list', { productID: device.productID, deviceName: device.deviceName }); // 查询始终带完整设备身份。
        if (generation !== this.generation || this.stopped) return false;
        listOf(response).forEach(record => { this._apply(deviceKey, record.dataID, record.value, record.timestamp); });
        this.records[deviceKey].error = false; this._notify(); return true;
      } catch (_) {
        if (generation === this.generation && !this.stopped) { this.records[deviceKey].error = true; this._notify(); }
        return false;
      }
    });
  };
  /** 查询选中设备指标的近一天历史；成功缓存一分钟，失败可重试，最多二十页限制单次工作。 */
  UrSceneProvider.prototype.fetchTrend = function (id, field) {
    field = field || 'p';
    const deviceKey = this.bindings[id], device = this.devices[deviceKey], config = this.config.fields[field], record = this.records[deviceKey], generation = this.generation; // 选中指标上下文。
    if (!device || !config || !config.identifier || !this.sdk || !this.sdk.callApi) return Promise.resolve(false);
    if (Date.now() - (record.historyAt[field] || 0) < 60000) return Promise.resolve(true);
    return this._request(deviceKey + ':history:' + field, async () => {
      const now = Date.now(), points = []; // 一次请求使用固定时间窗。
      record.trendStatus[field] = 'loading'; this._notify();
      try {
        for (let page = 1; page <= 20; page++) {
          if (generation !== this.generation || this.stopped) return false;
          const response = await this.sdk.callApi('/api/v1/things/device/msg/property-log/get-list', { productID: device.productID, deviceName: device.deviceName, dataID: config.identifier, timeStart: String(now - 86400000 - this.config.staleMs), timeEnd: String(now), order: 2, page: { page, size: 200 } }); // 精确设备查询，不访问全产品历史。
          if (generation !== this.generation || this.stopped) return false;
          const list = listOf(response); // 单页历史记录。
          list.forEach(item => {
            const value = this._convert(field, item.value), ts = timestamp(item.timestamp); // 历史与实时使用相同换算。
            if (typeof value === 'number' && ts && ts <= now) points.push({ t: ts, v: value });
          });
          if (list.length < 200 || page * 200 >= Number(dataOf(response).total)) break;
        }
        this._merge(record, field, points);
        record.trendStatus[field] = (record.history[field] || []).length ? 'ready' : 'empty';
        record.historyAt[field] = now; this._notify(); return true;
      } catch (_) {
        if (generation === this.generation && !this.stopped) { record.trendStatus[field] = 'error'; this._notify(); }
        return false;
      }
    });
  };
  /** 只有有效且明确的开关字段才显示合闸或分闸；其他读数不能推断开关状态。 */
  UrSceneProvider.switchState = function (state) {
    if (!state || state.status !== 'ready' || typeof state.on !== 'boolean') return 'unknown';
    return state.on ? 'on' : 'off';
  };
  /** 供 UI 获取标准趋势指标；不暴露生产配置。 */
  UrSceneProvider.TREND_METRICS = METRICS;
  window.PowerProviders = { UrSceneProvider };
})();
