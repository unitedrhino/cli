/*!
 * ur-scene 协议 v1 场景页接入 SDK（scene-sdk.js）
 *
 * 文件说明：自研 3D 场景页（如 AI 生成的 Three.js 单页）通过本 SDK 接入大屏
 * 「3D嵌入场景」组件（SceneEmbed）：上报可绑定锚点、接收绑定表与实时属性值、
 * 经桥接 RPC 调用平台 API、订阅项目级实时频道。
 *
 * 用法：<script src="scene-sdk.js"></script>
 *       const sdk = window.UrSceneSDK
 *       sdk.ready()                                  // 场景加载完成后上报（可靠握手：收到绑定表前每 1s 重发）
 *       sdk.reportAnchors([{ path, name, desc }])    // 上报可绑定点位
 *       sdk.onInit((bindings) => { ... })            // 收到大屏下发的绑定表
 *       sdk.onData(({ productID, deviceName, identifier, value, ts }) => { ... })
 *       sdk.onSelect((path) => { ... })              // 编辑态点位选中联动
 *       const res = await sdk.callApi('/api/v1/things/product/info/get-list', { page: { page: 1, size: 5 } })
 *       const unsub = sdk.subscribe('prop', (payload) => { ... })
 *
 * 安全说明：场景页在 sandbox="allow-scripts"（无 allow-same-origin）iframe 中运行，
 * 源不透明、无法读父页内容；所有平台请求均由大屏侧带登录态转发，凭证不出父页面。
 *
 * 握手说明（v1.1）：ready() 只发一次可能因竞态丢失（大屏组件重挂载、
 * 父页 message 监听注册晚于 iframe 首条消息等），故 ready() 发出后未收到
 * ur-scene:init 前按 1s 间隔重发（上限 60 次），收到 init 即停并触发 onInit 回调。
 * 消息格式与 v1 完全兼容：旧版大屏组件收到重发的 READY 会重复回发绑定表，无害。
 */
;(function (global) {
  'use strict'

  /** 协议消息类型常量（与大屏侧 scene-bridge.ts 保持一致） */
  var SCENE_MSG = {
    READY: 'ur-scene:ready',
    ANCHORS: 'ur-scene:anchors',
    INIT: 'ur-scene:init',
    DATA: 'ur-scene:data',
    API: 'ur-scene:api',
    API_RES: 'ur-scene:api:res',
    SUBSCRIBE: 'ur-scene:subscribe',
    WS: 'ur-scene:ws',
    SELECT: 'ur-scene:select'
  }

  /** 下行消息类型白名单（其余消息直接丢弃） */
  var ALLOWED_TYPES = {}
  ALLOWED_TYPES[SCENE_MSG.INIT] = true
  ALLOWED_TYPES[SCENE_MSG.DATA] = true
  ALLOWED_TYPES[SCENE_MSG.API_RES] = true
  ALLOWED_TYPES[SCENE_MSG.WS] = true
  ALLOWED_TYPES[SCENE_MSG.SELECT] = true

  /** 数据回调表 */
  var dataCallbacks = []
  /** 绑定表到达回调表 */
  var initCallbacks = []
  /** 选中联动回调表 */
  var selectCallbacks = []
  /** API 转发挂起的 Promise 表（id → { resolve, reject, timer }） */
  var pendingApis = {}
  /** WS 订阅回调表（订阅 id → 回调） */
  var wsSubscriptions = {}
  var apiSeq = 0
  var subSeq = 0
  var listenerBound = false
  /** ready 重发定时器（收到 ur-scene:init 前周期重发，收到即清除） */
  var readyRetryTimer = null
  /** ready 重发次数上限（1s × 60 = 最长 1 分钟，覆盖大屏侧任意慢启动） */
  var READY_RETRY_MAX = 60

  /** 向大屏 postMessage（sandbox 不透明源无法读取父页 origin，只能以 '*' 为目标） */
  function post(message) {
    if (global.parent && global.parent !== global) {
      global.parent.postMessage(message, '*')
    }
  }

  /** 超时清理挂起的 API 请求，避免场景页永久等待 */
  function rejectPending(id, msg) {
    var pending = pendingApis[id]
    if (!pending) return
    delete pendingApis[id]
    clearTimeout(pending.timer)
    pending.reject(new Error(msg))
  }

  /** 收到绑定表：停掉 ready 重发并触发 onInit 回调 */
  function handleInit(bindings) {
    if (readyRetryTimer) {
      clearInterval(readyRetryTimer)
      readyRetryTimer = null
    }
    for (var i = 0; i < initCallbacks.length; i++) {
      try {
        initCallbacks[i](bindings)
      } catch (err) {
        console.error('[UrSceneSDK] onInit 回调异常:', err)
      }
    }
  }

  /** 绑定 message 监听（幂等） */
  function bindListener() {
    if (listenerBound) return
    listenerBound = true
    global.addEventListener('message', function (event) {
      var data = event.data
      if (!data || typeof data !== 'object' || Array.isArray(data)) return
      var type = data.type
      if (typeof type !== 'string' || !ALLOWED_TYPES[type]) return

      if (type === SCENE_MSG.INIT) {
        handleInit(Array.isArray(data.bindings) ? data.bindings : [])
      } else if (type === SCENE_MSG.DATA) {
        for (var i = 0; i < dataCallbacks.length; i++) {
          try {
            dataCallbacks[i](data)
          } catch (err) {
            console.error('[UrSceneSDK] onData 回调异常:', err)
          }
        }
      } else if (type === SCENE_MSG.API_RES) {
        var pending = pendingApis[data.id]
        if (!pending) return
        delete pendingApis[data.id]
        clearTimeout(pending.timer)
        // code >= 400 视为业务/转发失败，走 reject 便于调用方统一 catch
        if (typeof data.code === 'number' && data.code >= 400) {
          pending.reject(new Error(data.msg || 'API 调用失败 code=' + data.code))
        } else {
          pending.resolve({ code: data.code, data: data.data, msg: data.msg })
        }
      } else if (type === SCENE_MSG.WS) {
        var callback = wsSubscriptions[data.id]
        if (callback) {
          try {
            callback(data.payload)
          } catch (err) {
            console.error('[UrSceneSDK] subscribe 回调异常:', err)
          }
        }
      } else if (type === SCENE_MSG.SELECT) {
        for (var j = 0; j < selectCallbacks.length; j++) {
          try {
            selectCallbacks[j](data.path || '')
          } catch (err) {
            console.error('[UrSceneSDK] onSelect 回调异常:', err)
          }
        }
      }
    })
  }

  /** 场景页 SDK 导出 */
  var UrSceneSDK = {
    SCENE_MSG: SCENE_MSG,

    /**
     * 场景页加载完成后调用：通知大屏可接收绑定表（大屏会回发 ur-scene:init）
     * 可靠握手：发出后未收到 ur-scene:init 前每 1s 重发（上限 60 次），
     * 覆盖大屏组件重挂载、父页监听注册晚于首条消息等竞态；收到 init 即停。
     * 旧版大屏组件收到重发的 READY 会重复回发绑定表，属协议兼容内的无害行为。
     */
    ready: function () {
      bindListener()
      var retries = 0
      post({ type: SCENE_MSG.READY })
      if (readyRetryTimer) clearInterval(readyRetryTimer)
      readyRetryTimer = setInterval(function () {
        if (++retries > READY_RETRY_MAX) {
          clearInterval(readyRetryTimer)
          readyRetryTimer = null
          console.warn('[UrSceneSDK] ready 重发达上限仍未收到绑定表，请检查大屏组件配置')
          return
        }
        post({ type: SCENE_MSG.READY })
      }, 1000)
    },

    /**
     * 注册绑定表到达回调（大屏回发 ur-scene:init 时触发）
     * @param {(bindings: Array<{ path: string, productID: string, deviceName: string, identifier: string }>) => void} callback 绑定表回调
     * @returns {function} 注销函数
     */
    onInit: function (callback) {
      bindListener()
      if (typeof callback !== 'function') return function () {}
      initCallbacks.push(callback)
      return function () {
        var index = initCallbacks.indexOf(callback)
        if (index >= 0) initCallbacks.splice(index, 1)
      }
    },

    /**
     * 上报可绑定锚点列表
     * @param {Array<{ path: string, name: string, desc?: string }>} anchors 锚点列表（path 需唯一）
     */
    reportAnchors: function (anchors) {
      if (!Array.isArray(anchors) || anchors.length === 0) {
        console.warn('[UrSceneSDK] reportAnchors 需要非空数组')
        return
      }
      post({ type: SCENE_MSG.ANCHORS, anchors: anchors })
    },

    /**
     * 注册实时属性值回调（每条 ur-scene:data 触发）
     * @param {(payload: { productID: string, deviceName: string, identifier: string, value: *, ts: number }) => void} callback
     * @returns {function} 注销函数
     */
    onData: function (callback) {
      bindListener()
      if (typeof callback !== 'function') return function () {}
      dataCallbacks.push(callback)
      return function () {
        var index = dataCallbacks.indexOf(callback)
        if (index >= 0) dataCallbacks.splice(index, 1)
      }
    },

    /**
     * 注册编辑态点位选中联动回调（大屏绑定面板选中锚点时触发）
     * @param {(path: string) => void} callback
     * @returns {function} 注销函数
     */
    onSelect: function (callback) {
      bindListener()
      if (typeof callback !== 'function') return function () {}
      selectCallbacks.push(callback)
      return function () {
        var index = selectCallbacks.indexOf(callback)
        if (index >= 0) selectCallbacks.splice(index, 1)
      }
    },

    /**
     * 经大屏代调平台 API（仅允许 /api/v1/ 前缀，仅 POST，凭证不出父页面）
     * @param {string} path 平台 API 路径
     * @param {*} body 请求体
     * @param {{ timeout?: number }} options 可选项（超时毫秒，默认 15s）
     * @returns {Promise<{ code: number, data: *, msg: string }>} 平台统一响应包
     */
    callApi: function (path, body, options) {
      bindListener()
      var timeout = (options && options.timeout) || 15000
      return new Promise(function (resolve, reject) {
        var id = 'api-' + Date.now() + '-' + ++apiSeq
        var timer = setTimeout(function () {
          rejectPending(id, 'API 调用超时: ' + path)
        }, timeout)
        pendingApis[id] = { resolve: resolve, reject: reject, timer: timer }
        post({ type: SCENE_MSG.API, id: id, path: path, body: body })
      })
    },

    /**
     * 订阅项目级实时频道（最小实现：'prop'=属性上报，'conn'=设备上下线）
     * @param {'prop'|'conn'} channel 频道
     * @param {(payload: *) => void} callback 命中消息回调
     * @returns {function} 退订函数
     */
    subscribe: function (channel, callback) {
      bindListener()
      if (typeof callback !== 'function') return function () {}
      var id = 'sub-' + Date.now() + '-' + ++subSeq
      wsSubscriptions[id] = callback
      post({ type: SCENE_MSG.SUBSCRIBE, id: id, channel: channel })
      return function () {
        delete wsSubscriptions[id]
      }
    }
  }

  global.UrSceneSDK = UrSceneSDK
})(typeof window !== 'undefined' ? window : this)
