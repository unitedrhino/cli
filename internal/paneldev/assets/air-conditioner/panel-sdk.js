/* 文件说明：无依赖的面板网页 SDK；Web/App 共用属性控制、订阅、运行期认证与平台接口调用合同。 */
(function createPanelSDK(global) {
  /** 协议版本限定当前面板消息命名空间。 */
  const protocol = 'ur-panel/v1';
  /** pending 保存关联请求；订阅回调仅留在当前网页。 */
  const pending = new Map();
  const subscriptions = new Map();
  /** port/native 是两种互斥的本地传输实现。 */
  let closed = false;
  let port;
  let readyTimer;
  let retryTimer;
  let sequence = 0;
  let sessionID;
  /** nativeBridge 由可信 App 壳装配，不能因此跳过原生层权限校验。 */
  const nativeBridge = global.__UR_PANEL_NATIVE__;
  /** readyResolve/readyReject 完成唯一初始化握手。 */
  let readyReject, readyResolve;
  const initialized = new Promise((resolve, reject) => {
    readyResolve = resolve;
    readyReject = reject;
  });
  initialized.catch(() => {});
  /** failure 构造可供调用方识别的语义异常。 */
  function failure(code, message) {
    return Object.assign(new Error(message), { code });
  }
  /** send 只通过当前本机桥发消息，浏览器消息本身不经过服务器。 */
  function send(message) {
    if (nativeBridge) nativeBridge.postMessage(message);
    else if (port) port.postMessage(message);
  }
  /** receive 只把响应投递给当前会话的关联请求，不向其他窗口广播认证或设备数据。 */
  function receive(message) {
    if (!message || message.protocol !== protocol || closed) return;
    if (message.type === 'init' && !sessionID && typeof message.sessionID === 'string') {
      sessionID = message.sessionID;
      clearInterval(retryTimer);
      clearTimeout(readyTimer);
      readyResolve();
      return;
    }
    if (!sessionID || message.sessionID !== sessionID) return;
    switch (message.type) {
      case 'close': {
        close();
        break;
      }
      case 'event': {
        for (const subscription of subscriptions.values())
          if (!subscription.hostID && subscription.topics.includes(message.topic)) {
            if (subscription.early.length < 100) subscription.early.push(message);
          } else if (subscription.hostID === message.subscriptionID && subscription.topics.includes(message.topic)) {
            try {
              subscription.callback({ topic: message.topic, data: message.data, timestamp: message.timestamp });
            } catch (error) {
              console.error('[urPanel] 订阅回调失败', error);
            }
          }

        break;
      }
      case 'response': {
        const request = pending.get(message.id);
        if (!request) return;
        pending.delete(message.id);
        clearTimeout(request.timer);
        if (message.ok) request.resolve(message.result);
        else request.reject(failure(message.error?.code || 'INVALID_PARAMS', message.error?.message || '面板调用失败'));

        break;
      }
    }
  }
  /** receiveInit 仅信任父窗口，随后只使用转移的专用端口。 */
  function receiveInit(event) {
    if (event.source !== global.parent || event.data?.protocol !== protocol || event.data.type !== 'init' || sessionID || !event.ports?.[0]) return;
    port = event.ports[0];
    port.addEventListener('message', (event) => receive(event.data));
    port.start();
    receive(event.data);
  }
  /** ready 重发只用于未完成的初始化，不重放任何控制请求。 */
  function ready() {
    const message = { protocol, type: 'ready' };
    if (nativeBridge) nativeBridge.postMessage(message);
    else if (global.parent !== global) global.parent.postMessage(message, '*');
  }
  /** call 等初始化后发一个请求；超时不会重发或改变设备通道。 */
  async function call(method, params = {}) {
    await initialized;
    if (closed) throw failure('SESSION_CLOSED', '面板会话已关闭');
    const id = `request-${++sequence}`;
    return new Promise((resolve, reject) => {
      const timer = setTimeout(() => {
        pending.delete(id);
        reject(failure('TIMEOUT', '面板调用未获得结果，请核对设备状态'));
      }, 10_000);
      pending.set(id, { resolve, reject, timer });
      try {
        send({ protocol, type: 'request', sessionID, id, method, params });
      } catch (error) {
        clearTimeout(timer);
        pending.delete(id);
        reject(error);
      }
    });
  }
  /** close 释放当前生命周期全部请求，旧回包不能进入新面板。 */
  function close() {
    if (closed) return;
    closed = true;
    clearInterval(retryTimer);
    clearTimeout(readyTimer);
    readyReject(failure('SESSION_CLOSED', '面板会话已关闭'));
    for (const request of pending.values()) {
      clearTimeout(request.timer);
      request.reject(failure('SESSION_CLOSED', '面板会话已关闭'));
    }
    pending.clear();
    subscriptions.clear();
    port?.close();
    global.removeEventListener('message', receiveInit);
  }
  /** sdk 保留单设备语义方法；直接调用现有平台接口时由宿主补齐当前登录上下文。 */
  const sdk = Object.freeze({
    getContext: () => call('getContext'),
    getAuth: () => call('getAuth'),
    request: (params) => call('request', params),
    getThingModel: () => call('getThingModel'),
    readProperties: (params) => call('readProperties', params),
    writeProperties: (params) => call('writeProperties', params),
    invokeAction: (params) => call('invokeAction', params),
    subscribe: async (params, callback) => {
      if (typeof callback !== 'function' || !Array.isArray(params?.topics)) throw failure('INVALID_PARAMS', '订阅需要主题和回调');
      const localID = `subscription-${++sequence}`;
      subscriptions.set(localID, { topics: [...params.topics], callback, early: [] });
      try {
        const result = await call('subscribe', { topics: params.topics });
        const subscription = subscriptions.get(localID);
        subscription.hostID = result.subscriptionID;
        for (const event of subscription.early) {
          if (event.subscriptionID === result.subscriptionID) {
            try {
              callback({ topic: event.topic, data: event.data, timestamp: event.timestamp });
            } catch (error) {
              console.error('[urPanel] 订阅回调失败', error);
            }
          }
        }
        subscription.early.length = 0;
        let removed = false;
        return () => {
          if (removed) return;
          removed = true;
          subscriptions.delete(localID);
          call('unsubscribe', { subscriptionID: result.subscriptionID }).catch(() => {});
        };
      } catch (error) {
        subscriptions.delete(localID);
        throw error;
      }
    },
  });
  Object.defineProperty(global, 'urPanel', { value: sdk, configurable: false, writable: false });
  if (nativeBridge) Object.defineProperty(global, '__urPanelReceive', { value: receive });
  else global.addEventListener('message', receiveInit);
  global.addEventListener('pagehide', close, { once: true });
  readyTimer = setTimeout(() => {
    clearInterval(retryTimer);
    readyReject(failure('TIMEOUT', '面板宿主初始化超时'));
  }, 5000);
  retryTimer = setInterval(ready, 1000);
  ready();
})(window);
