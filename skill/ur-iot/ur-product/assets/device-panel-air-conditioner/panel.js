/* 文件说明：共用设备面板交互；保留未知/只读/离线语义，受理与实际回执分别展示，不把模拟或缓存当实时数据。 */
(() => {
  /** state 保存当前宿主可用字段、设备上报值和未完成命令；禁止以界面默认值伪造设备状态。 */
  const state = {
    fields: new Map(),
    values: {},
    ready: false,
    writable: false,
    busy: false,
    pending: null,
    early: new Map(),
    timer: null,
  };
  /** $ 只查找面板自己的元素，不读取父窗口。 */
  const $ = (selector) => document.querySelector(selector);
  /** text 以纯文本更新可选控件，防止设备文本作为 HTML 插入。 */
  const text = (selector, value) => {
    const element = $(selector);
    if (element) element.textContent = value;
  };
  /** modes 映射示例枚举与用户文案；接入产品时须与真实模型同时修改。 */
  const modes = { 0: '制冷', 1: '制热', 2: '除湿', 3: '自动' };
  /** labels 只承认最终回执状态，受理不会映射成成功。 */
  const labels = {
    succeeded: '设备已确认，设置成功',
    failed: '设备执行失败，请检查后重试',
    timeout: '等待回执超时，执行结果未知',
  };
  /** feedback 将用户可理解的结果放入固定状态区，技术细节保留给模拟宿主。 */
  function feedback(message, error = false) {
    text('#feedback', message);
    $('#feedback').dataset.state = error ? 'error' : 'normal';
  }
  /** value 仅采用质量有效的上报值；未知值保持 null，不补造温度或关闭状态。 */
  function value(id) {
    const field = state.values[id];
    return !field ||
      ['UNKNOWN', 'BAD', 'INVALID'].includes(String(field.quality || 'UNKNOWN').toUpperCase())
      ? null
      : field.value;
  }
  /** allowed 以实际模型和联网能力共同决定控件权限，关闭后仅允许电源操作。 */
  function allowed(id) {
    const field = state.fields.get(id);
    return (
      state.ready &&
      state.writable &&
      !state.busy &&
      field?.available !== false &&
      field?.mode === 'rw' &&
      (id === 'power' || value('power') === true)
    );
  }
  /** render 从真实状态渲染统一选中、禁用及边界；温控范围从模型获取。 */
  function render() {
    const power = value('power'),
      temperature = value('targetTemperature'),
      mode = value('mode');
    text('#power-label', power === null ? '状态未知' : power ? '空调已开启' : '设备已关闭');
    if (!$('#temperature'))
      text('#power-label', power === null ? '状态未知' : power ? '设备已开启' : '设备已关闭');
    $('#power').setAttribute('aria-pressed', String(power === true));
    $('#power').disabled = !allowed('power') || typeof power !== 'boolean';
    text('#temperature', typeof temperature === 'number' ? temperature : '—');
    const arc = $('.dial .arc');
    if (arc) {
      const range = state.fields.get('targetTemperature'),
        min = Number(range?.min ?? 16),
        max = Number(range?.max ?? 30);
      const progress =
        typeof temperature === 'number' && max > min
          ? Math.max(0, Math.min(1, (temperature - min) / (max - min)))
          : 0;
      arc.style.strokeDasharray = `${Math.max(12, progress * 407)} 600`;
      arc.style.opacity = power === true ? '0.9' : '0.25';
    }
    text('#mode-label', modes[mode] || '模式未知');
    text(
      '#climate-caption',
      power === null
        ? '等待设备状态'
        : !power
          ? '开机后可调节温度'
          : modes[mode]
            ? modes[mode] + '运行中'
            : '运行中',
    );
    const ambient = value('currentTemperature');
    text(
      '#ambient-label',
      typeof ambient === 'number' ? `室内温度 · ${ambient} °C` : '室内温度 · 暂无数据',
    );
    const field = state.fields.get('targetTemperature'),
      min = Number(field?.min ?? 16),
      max = Number(field?.max ?? 30);
    text('#temperature-range', `${min} — ${max} °C`);
    document.querySelectorAll('[data-step]').forEach((button) => {
      const next =
        typeof temperature === 'number'
          ? temperature + Number(button.dataset.step) * Number(field?.step || 1)
          : NaN;
      button.disabled =
        !allowed('targetTemperature') || !Number.isFinite(next) || next < min || next > max;
    });
    document.querySelectorAll('[data-property]').forEach((button) => {
      button.disabled = !allowed(button.dataset.property);
      button.setAttribute(
        'aria-pressed',
        String(String(value(button.dataset.property)) === button.dataset.value),
      );
    });
  }
  /** finish 只结束当前命令；超时不回放控制，后来的回执也不误覆盖新命令。 */
  function finish(result) {
    if (result.commandID !== state.pending || !labels[result.status]) return;
    clearTimeout(state.timer);
    state.pending = null;
    state.busy = false;
    feedback(labels[result.status], result.status !== 'succeeded');
    render();
  }
  /** send 单次提交后等待实际设备回执，等待期间锁定控件，拒绝乐观伪造执行成功。 */
  async function send(values) {
    if (state.busy || Object.keys(values).some((id) => !allowed(id))) return;
    state.busy = true;
    render();
    feedback('正在发送设置…');
    try {
      const accepted = await urPanel.writeProperties({ values });
      state.pending = accepted.commandID;
      feedback('设置已受理，等待设备确认…');
      if (labels[accepted.status]) finish(accepted);
      else if (state.early.has(accepted.commandID)) finish(state.early.get(accepted.commandID));
      if (state.pending)
        state.timer = setTimeout(
          () => finish({ commandID: accepted.commandID, status: 'timeout' }),
          32000,
        );
      state.early.clear();
    } catch (error) {
      state.busy = false;
      feedback(error.message || '设置发送失败', true);
      render();
    }
  }
  $('#power').addEventListener('click', () => send({ power: !value('power') }));
  document.querySelectorAll('[data-step]').forEach((button) =>
    button.addEventListener('click', () =>
      send({
        targetTemperature:
          value('targetTemperature') +
          Number(button.dataset.step) * Number(state.fields.get('targetTemperature')?.step || 1),
      }),
    ),
  );
  document
    .querySelectorAll('[data-property]')
    .forEach((button) =>
      button.addEventListener('click', () =>
        send({
          [button.dataset.property]:
            state.fields.get(button.dataset.property)?.type === 'enum'
              ? Number(button.dataset.value)
              : button.dataset.value,
        }),
      ),
    );
  /** start 先订阅再读值，避免初始化期间漏报；卸载时销毁当前订阅和等待定时器。 */
  async function start() {
    try {
      const context = await urPanel.getContext(),
        model = await urPanel.getThingModel();
      state.fields = new Map(model.properties.map((field) => [field.identifier, field]));
      state.writable =
        context.mode !== 'offline' && context.capabilities?.cloud?.available === true;
      const mock = context.mode === 'mock';
      text('#device-name', mock ? '客厅 · 模拟设备' : context.deviceName || '设备控制');
      text('#connection-label', mock ? '模拟运行' : state.writable ? '云端已连接' : '离线 · 只读');
      $('#connection').dataset.connected = String(state.writable);
      text(
        '#data-note',
        mock
          ? '模拟数据 · 不控制真实设备'
          : state.writable
            ? '状态由设备上报'
            : '缓存快照 · 非实时数据',
      );
      if (!state.writable) text('#control-note', '当前仅供查看，联网后可控制设备。');
      const unsubscribe = await urPanel.subscribe(
        { topics: ['property', 'command', 'connection'] },
        (event) => {
          if (event.topic === 'property') {
            Object.assign(state.values, event.data);
            render();
          }
          if (event.topic === 'command') {
            if (event.data.commandID === state.pending) finish(event.data);
            else if (state.busy && !state.pending && state.early.size < 20)
              state.early.set(event.data.commandID, event.data);
          }
          if (event.topic === 'connection') {
            // 通道断开立即锁定控制，恢复后允许重新操作，不重发旧命令。
            state.writable = event.data.cloud?.connected === true;
            if (event.data.deviceOnline === false) state.writable = false;
            text(
              '#connection-label',
              state.writable
                ? mock
                  ? '模拟运行'
                  : '设备已连接'
                : event.data.deviceOnline === false
                  ? '设备离线'
                  : '连接已断开',
            );
            $('#connection').dataset.connected = String(state.writable);
            render();
          }
        },
      );
      window.addEventListener(
        'pagehide',
        () => {
          unsubscribe();
          clearTimeout(state.timer);
        },
        { once: true },
      );
      Object.assign(
        state.values,
        await urPanel.readProperties({
          identifiers: model.properties
            .filter((field) => field.available !== false)
            .map((field) => field.identifier),
        }),
      );
      state.ready = true;
      render();
      feedback(state.writable ? '设置随设备状态同步' : '离线模式，仅展示上次缓存状态');
    } catch (error) {
      state.ready = false;
      render();
      feedback(error.message || '无法读取设备状态，请重新打开面板', true);
      text('#connection-label', '连接不可用');
    }
  }
  start();
})();
