/* ============================================================
 * main.js — 装配入口
 * 初始化场景 → 数据订阅 → UI（设备树 / 信息面板 / 工具栏 / 统计）
 * ============================================================ */

(function () {
  const registry = window.buildMeterRegistry();
  const registryMap = {};
  registry.forEach(m => registryMap[m.id] = m);

  const container = document.getElementById('viewport');
  const api = window.StationScene.build(container, registry);
  const inter = window.setupInteraction(api, {
    onSelect: id => openPanel(id, false),
    onDeselect: () => closePanel(),
    onHover: (id, e) => {
      if (id) { showTooltip(id, e); } else hideTooltip();
    },
    onHoverMove: e => moveTooltip(e)
  });

  /* ---------- 数据提供层：无绑定时保留空状态，宿主负责凭证与设备授权。 ---------- */
  const provider = new window.PowerProviders.UrSceneProvider(registry); // 真实数据桥。
  const srcTip = document.querySelector('#info-panel .src'); // 状态说明位置。
  const SRC_TEXT = { pending: '正在连接数据宿主…', real: '数据来源：设备物模型 · 双击电表聚焦', unbound: '未连接或未配置电表绑定' }; // 明确区分未绑定与无读数。
  provider.onModeChange = mode => { if (srcTip) srcTip.textContent = SRC_TEXT[mode] || SRC_TEXT.unbound; };
  provider.onModeChange(provider.mode);
  let latest = provider.states;
  provider.subscribe(states => {
    latest = states;
    api.refreshLcds(states, registryMap);
    updateSummary();
    updateTreeStates();
    if (currentId) fillPanel(currentId);
  });


  /* ---------- 顶部统计 ---------- */
  const $sum = {
    load: document.getElementById('sum-load'),
    online: document.getElementById('sum-online'),
    kwh: document.getElementById('sum-kwh'),
    clock: document.getElementById('sum-clock')
  };
  function updateSummary() {
    const load = provider.summary('p'), kwh = provider.summary('kwh'); // 仅汇总显式核验的进线关系。
    const valid = registry.filter(meter => latest[meter.id] && latest[meter.id].status === 'ready').length; // 不把有绑定等同在线。
    $sum.load.textContent = load === null ? (window.STATION_DATA.summary.verified ? '-- kW' : '未配置汇总') : load.toFixed(2) + ' kW';
    $sum.online.textContent = valid + ' / ' + registry.length;
    $sum.kwh.textContent = kwh === null ? '-- kWh' : kwh.toFixed(1) + ' kWh';
  }
  const clockTimer = setInterval(() => { // 时钟计时器随页面销毁。
    $sum.clock.textContent = new Date().toLocaleTimeString('zh-CN', { hour12: false });
  }, 1000);

  /* ---------- 左侧设备树 ---------- */
  const tree = document.getElementById('device-tree');
  document.getElementById('tree-count').textContent = registry.length + ' 台';
  function buildTree() {
    let html = '';
    html += group('高压区', registry.filter(m => m.kind === 'hv'));
    html += group('变压器', registry.filter(m => m.kind === 'trans'));
    window.LV_SECTIONS.forEach(sec => {
      html += '<div class="tree-sec">' + sec.name + '</div>';
      sec.cabinets.forEach(cab => {
        const meters = registry.filter(m => m.cabinet === cab.id);
        html += group2(cab.name + ' (' + cab.id + ')', meters);
      });
    });
    tree.innerHTML = html;

    tree.addEventListener('click', e => {
      const item = e.target.closest('[data-mid]');
      const head = e.target.closest('.tree-group-head');
      if (item) {
        const id = item.getAttribute('data-mid');
        tree.querySelectorAll('.tree-item.sel').forEach(x => x.classList.remove('sel'));
        item.classList.add('sel');
        inter.select(id, true);
        openPanel(id, true);
      } else if (head) {
        head.parentElement.classList.toggle('collapsed');
      }
    });
  }
  function group(title, meters) {
    if (!meters.length) return '';
    return group2(title, meters, 'open');
  }
  function group2(title, meters) {
    let html = '<div class="tree-group"><div class="tree-group-head">' + title +
      ' <span class="tree-count">' + meters.length + '</span></div><div class="tree-body">';
    meters.forEach(m => {
      html += '<div class="tree-item" data-mid="' + m.id + '"><span class="dot" data-state="unknown" title="状态未知"></span>' +
        m.name + '<span class="tree-id">' + m.id + '</span></div>';
    });
    return html + '</div></div>';
  }
  /** 树节点只根据实际有效开关字段着色；未绑定、过期和缺开关字段保持灰色。 */
  function updateTreeStates() {
    tree.querySelectorAll('[data-mid]').forEach(item => {
      const dot = item.querySelector('.dot'), state = window.PowerProviders.UrSceneProvider.switchState(latest[item.dataset.mid]); // 当前节点开关展示态。
      dot.dataset.state = state;
      dot.title = state === 'on' ? '合闸' : state === 'off' ? '分闸' : '状态未知';
      dot.setAttribute('aria-label', dot.title);
    });
  }
  buildTree();
  updateTreeStates();

  /* ---------- 右侧信息面板 ---------- */
  const panel = document.getElementById('info-panel');
  const trendCv = document.getElementById('trend');
  let currentId = null;
  const trendStatus = document.getElementById('trend-status'); // 历史请求的可访问空状态。
  const trendRetry = document.getElementById('trend-retry'); // 失败后的手动重试，不新增轮询。
  trendRetry.addEventListener('click', () => {
    if (currentId && typeof provider.fetchTrend === 'function') provider.fetchTrend(currentId, trendMetric);
  });

  /** 趋势指标与物模型映射由配置和数据桥定义，不再维护演示分支。 */
  const TREND_METRICS = window.PowerProviders.UrSceneProvider.TREND_METRICS;
  let trendMetric = 'p';

  /* 趋势指标切换页签：切换即重绘，并按需回填该指标近 24h 历史 */
  const trendTabs = document.getElementById('trend-tabs');
  TREND_METRICS.forEach(mt => {
    const b = document.createElement('button');
    b.className = 'trend-tab' + (mt.field === trendMetric ? ' active' : '');
    b.textContent = mt.unit ? mt.label + ' ' + mt.unit : mt.label;
    b.addEventListener('click', () => {
      if (trendMetric === mt.field) return;
      trendMetric = mt.field;
      trendTabs.querySelectorAll('.trend-tab').forEach(x => x.classList.toggle('active', x === b));
      if (!currentId) return;
      drawTrend(latest[currentId]);
      if (typeof provider.fetchTrend === 'function' && provider.mode === 'real') {
        provider.fetchTrend(currentId, mt.field);
      }
    });
    trendTabs.appendChild(b);
  });

  function openPanel(id, fly) {
    currentId = id;
    panel.classList.add('open');
    fillPanel(id);
    refreshDetail(id);
    /* 趋势曲线按实际上报追加攒得太慢，打开面板时回填历史 */
    if (typeof provider.fetchTrend === 'function' && provider.mode === 'real') {
      provider.fetchTrend(id, trendMetric);
    }
  }
  function closePanel() {
    currentId = null;
    panel.classList.remove('open');
  }

  /* 面板打开时按需补拉全量字段（UrSceneProvider 真实模式下经大屏代调 API），
   * 面板停留期间按配置刷新；无绑定时不发送查询 */
  function refreshDetail(id) {
    if (typeof provider.fetchDetail === 'function' && provider.mode === 'real') {
      provider.fetchDetail(id);
    }
  }
  const detailRefreshTimer = setInterval(() => { if (currentId) refreshDetail(currentId); }, window.STATION_DATA.pollMs);
  document.getElementById('panel-close').addEventListener('click', () => {
    inter.deselect(); closePanel();
  });

  function fmt(v, d, u) { return (v == null ? '--' : Number(v).toFixed(d == null ? 1 : d)) + (u || ''); }

  function fillPanel(id) {
    const m = registryMap[id];
    const s = latest[id];
    if (!m) return;
    document.getElementById('p-name').textContent = m.name;
    document.getElementById('p-id').textContent = m.id + ' · ' + m.meterType;
    const badge = document.getElementById('p-status');
    const statusText = { unbound: '未绑定', unconfigured: '物模型未配置', empty: '暂无数据', error: '连接失败', stale: '数据过期', ready: '数据有效' }; // 状态不猜测设备开关。
    badge.textContent = statusText[s.status] || '暂无数据';
    badge.className = 'badge ' + (s.status === 'ready' ? 'on' : 'unknown');

    const body = document.getElementById('p-body');
    body.innerHTML =
      rows([['所属柜体', m.cabinetName], ['所属区域', m.section],
            ['表计', m.meterType], ['回路编号', m.id]]) +
      rows3('相电压 (V)', [fmt(s && s.ua, 1), fmt(s && s.ub, 1), fmt(s && s.uc, 1)]) +
      rows3('电流 (A)', [fmt(s && s.ia, 1), fmt(s && s.ib, 1), fmt(s && s.ic, 1)]) +
      rows([['有功功率', fmt(s && s.p, 2, ' kW')], ['无功功率', fmt(s && s.q, 2, ' kvar')],
            ['功率因数', fmt(s && s.pf, 3)], ['有功电能', fmt(s && s.kwh, 1, ' kWh')],
            ['开关状态', s && s.on != null ? (s.on ? '合闸' : '分闸') : '--'],
            ['数据更新', s && s.ts ? new Date(s.ts).toLocaleTimeString('zh-CN', { hour12: false }) : '--']]);
    drawTrend(s);
  }

  function rows(pairs) {
    let h = '<div class="p-grid">';
    pairs.forEach(([k, v]) => {
      h += '<div class="p-k">' + k + '</div><div class="p-v">' + v + '</div>';
    });
    return h + '</div>';
  }
  function rows3(label, vals) {
    return '<div class="p-3row"><div class="p-3label">' + label + '</div><div class="p-3cells">' +
      vals.map(v => '<div class="p-cell">' + v + '</div>').join('') + '</div></div>';
  }
  /* 趋势指标刻度值缩写：位数过长时降精度，避免 Y 轴标签挤出绘图区 */
  function tickText(v, decimals) {
    const s = Number(v).toFixed(decimals == null ? 1 : decimals);
    return s.length > 7 ? Number(v).toFixed(0) : s;
  }
  function hhmm(t) {
    const d = new Date(t);
    const p = n => String(n).padStart(2, '0');
    return p(d.getHours()) + ':' + p(d.getMinutes());
  }

  function drawTrend(s) {
    const mt = TREND_METRICS.find(m => m.field === trendMetric) || TREND_METRICS[0];
    const c = trendCv.getContext('2d');
    const W = trendCv.width, H = trendCv.height;
    c.clearRect(0, 0, W, H);
    c.fillStyle = 'rgba(10,16,24,0.6)';
    c.fillRect(0, 0, W, H);
    c.strokeStyle = 'rgba(120,160,200,0.25)';
    c.strokeRect(0.5, 0.5, W - 1, H - 1);
    c.fillStyle = '#7fa8c9';
    c.font = '12px Consolas, monospace';
    c.textAlign = 'left';
    c.fillText('实时趋势: ' + mt.label + (mt.unit ? ' ' + mt.unit : ''), 6, 14);
    /* 请求状态与数据点数量分开：空成功响应结束加载，单点不伪造曲线。 */
    const series = s && s.trends ? (s.trends[trendMetric] || []) : [];
    const validPoints = series.filter(point => point.v !== null); // 缺口不计作有效历史点。
    if (validPoints.length < 2) {
      const status = s && s.trendStatus && s.trendStatus[trendMetric]; // 当前电表、当前指标的请求结果。
      const loading = status === 'loading', failed = status === 'error'; // 请求进行中与失败不推断为无数据。
      const unbound = s && s.status === 'unbound', unconfigured = !window.STATION_DATA.fields[trendMetric].identifier; // 无绑定和无指标映射分别说明。
      trendStatus.hidden = false;
      trendStatus.dataset.state = loading ? 'loading' : failed ? 'error' : validPoints.length ? 'single' : status === 'empty' ? 'empty' : 'waiting';
      document.getElementById('trend-status-title').textContent = unbound ? '未绑定设备' : unconfigured ? '未配置指标' : loading ? '正在加载历史数据' : failed ? '历史数据加载失败' : validPoints.length ? '仅有 1 条记录' : status === 'empty' ? '近24小时暂无数据' : '等待数据上报';
      document.getElementById('trend-status-hint').textContent = unbound ? '请先在宿主配置电表锚点绑定' : unconfigured ? '请核验并配置物模型标识符与单位' : loading ? '正在查询近24小时记录…' : failed ? '请稍后重试' : validPoints.length ? '收到更多采样后生成趋势' : '暂无' + mt.label + '记录，上报后自动显示';
      trendRetry.hidden = !failed;
      return;
    }
    trendStatus.hidden = true;
    const d = series; // 保留完整一天的时间网格。
    /* 绘图区四周留白：左 Y 轴刻度、下 X 轴时间、上标题行 */
    const padL = 44, padR = 8, padT = 22, padB = 18;
    const pw = W - padL - padR, ph = H - padT - padB;
    /* 按数据实际 min~max 自适应量程：负荷变化相对绝对值可能很小，
       0 基量程会把曲线压成顶部一条直线，看不出波动；恒值序列兜底 ±10% 量程 */
    const vals = d.filter(point => point.v !== null).map(point => point.v);
    const mn = Math.min(...vals), mx = Math.max(...vals);
    const span = (mx - mn) || Math.abs(mx) * 0.2 || 1;
    const lo = mn - span * 0.15, hi = mx + span * 0.15;
    const t0 = d[0].t, t1 = d[d.length - 1].t;
    const xOf = i => padL + pw * (t1 > t0 ? (d[i].t - t0) / (t1 - t0) : (d.length > 1 ? i / (d.length - 1) : 0));
    const yOf = v => padT + ph - ph * (v - lo) / (hi - lo);
    /* Y 轴：轴线 + 上/中/下三档刻度值与横向网格线 */
    c.strokeStyle = 'rgba(120,160,200,0.35)';
    c.lineWidth = 1;
    c.beginPath();
    c.moveTo(padL, padT); c.lineTo(padL, padT + ph); c.lineTo(padL + pw, padT + ph);
    c.stroke();
    c.font = '10px Consolas, monospace';
    c.fillStyle = 'rgba(127,168,201,0.85)';
    c.textAlign = 'right';
    [hi, (hi + lo) / 2, lo].forEach((v, i) => {
      const y = padT + ph * i / 2;
      c.fillText(tickText(v, mt.decimals), padL - 4, y + 3);
      c.strokeStyle = 'rgba(120,160,200,0.12)';
      c.beginPath(); c.moveTo(padL, y + 0.5); c.lineTo(padL + pw, y + 0.5); c.stroke();
    });
    /* X 轴：首/中/尾时间标签（HH:mm），首尾分别左/右对齐避免出界 */
    c.fillStyle = 'rgba(127,168,201,0.7)';
    const xi = [0, (d.length - 1) >> 1, d.length - 1];
    const aligns = ['left', 'center', 'right'];
    xi.forEach((di, k) => {
      const p = d[di];
      if (!p || !p.t) return;
      c.textAlign = aligns[k];
      const x = k === 0 ? padL : (k === 2 ? padL + pw : xOf(di));
      c.fillText(hhmm(p.t), x, padT + ph + 13);
    });
    c.textAlign = 'left';
    /* 曲线本体 */
    c.strokeStyle = '#2ec8ff';
    c.lineWidth = 1.8;
    c.beginPath();
    let drawing = false; // 缺失或过期采样之间不连线。
    d.forEach((p, i) => {
      if (p.v === null) { drawing = false; return; }
      const x = xOf(i), y = yOf(p.v); // 当前有效点位置。
      drawing ? c.lineTo(x, y) : c.moveTo(x, y);
      drawing = true;
    });
    c.stroke();
    /* 右上角标注最新值，便于读数 */
    c.fillStyle = '#7fa8c9';
    c.font = '12px Consolas, monospace';
    c.textAlign = 'right';
    c.fillText((d[d.length - 1].v === null ? '--' : d[d.length - 1].v.toFixed(mt.decimals)) + (mt.unit ? ' ' + mt.unit : ''), W - 6, 13);
    c.textAlign = 'left';
  }

  /* ---------- 悬停 tooltip ---------- */
  const tip = document.getElementById('tooltip');
  function showTooltip(id, e) {
    const m = registryMap[id];
    tip.innerHTML = '<b>' + m.name + '</b><span>' + m.id + ' · 点击查看详情</span>';
    tip.style.display = 'block';
    moveTooltip(e);
  }
  function moveTooltip(e) {
    tip.style.left = (e.clientX + 14) + 'px';
    tip.style.top = (e.clientY + 14) + 'px';
  }
  function hideTooltip() { tip.style.display = 'none'; }

  /* ---------- 工具栏 ---------- */
  document.querySelectorAll('[data-view]').forEach(btn => {
    btn.addEventListener('click', () => {
      document.querySelectorAll('[data-view]').forEach(b => b.classList.remove('active'));
      btn.classList.add('active');
      api.flyTo(api.views[btn.getAttribute('data-view')]);
    });
  });
  const autoBtn = document.getElementById('btn-auto');
  autoBtn.addEventListener('click', () => {
    api.controls.autoRotate = !api.controls.autoRotate;
    autoBtn.classList.toggle('active', api.controls.autoRotate);
  });

  /* ---------- 自动巡检（轮播选中电表） ---------- */
  const patrolBtn = document.getElementById('btn-patrol');
  let patrolTimer = null, patrolIdx = 0;
  const patrolList = registry;
  function patrolStep() {
    const m = patrolList[patrolIdx % patrolList.length];
    patrolIdx++;
    inter.select(m.id, true);
    openPanel(m.id, true);
    const item = tree.querySelector('[data-mid="' + m.id + '"]');
    if (item) {
      tree.querySelectorAll('.tree-item.sel').forEach(x => x.classList.remove('sel'));
      item.classList.add('sel');
      item.scrollIntoView({ block: 'nearest' });
    }
  }
  patrolBtn.addEventListener('click', () => {
    if (patrolTimer) {
      clearInterval(patrolTimer); patrolTimer = null;
      patrolBtn.classList.remove('active');
    } else {
      patrolStep();
      patrolTimer = setInterval(patrolStep, 6000);
      patrolBtn.classList.add('active');
    }
  });

  /* ---------- 数据来源说明 ---------- */
  document.getElementById('btn-about').addEventListener('click', () => {
    document.getElementById('about-modal').classList.add('open');
  });
  document.getElementById('about-close').addEventListener('click', () => {
    document.getElementById('about-modal').classList.remove('open');
  });

  /* iframe 卸载时释放新增图形资源；bfcache 暂存页面保持可恢复。 */
  window.addEventListener('pagehide', event => {
    if (event.persisted) return;
    clearInterval(clockTimer);
    clearInterval(detailRefreshTimer);
    clearInterval(patrolTimer);
    provider.stop();
    inter.dispose();
    api.stop();
  });

  updateSummary();
  api.refreshLcds(latest, registryMap);
  // 场景首帧先呈现，再初始化宿主握手，数据接口不会阻塞模型加载。
  requestAnimationFrame(() => setTimeout(() => provider.start(), 0));
  window.StationApp = { provider, api, inter, registry, openPanel }; // 浏览器验收入口，不持有任何凭证。
})();
