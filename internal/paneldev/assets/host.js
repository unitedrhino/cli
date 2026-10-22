/* 自动生成：来源 apps/web/packages/saas/src/device-panel/dev-host.ts。 */
"use strict";
(() => {
  // apps/web/packages/saas/src/device-panel/bridge.ts
  function bridgeError(code, message) {
    return Object.assign(new Error(message), { code });
  }
  function validateApiRequest(params) {
    requestObject(params);
    onlyKeys(params, ["path", "data"]);
    if (typeof params.path !== "string" || !/^\/api\/v1\/[\w-]+(?:\/[\w-]+)*$/.test(params.path))
      throw bridgeError("INVALID_PARAMS", "\u8BF7\u4F7F\u7528 /api/v1/ \u5F00\u5934\u7684\u5E73\u53F0\u63A5\u53E3\u8DEF\u5F84");
    if (params.data !== void 0) requestObject(params.data);
  }
  function requestObject(value) {
    if (!value || Object.prototype.toString.call(value) !== "[object Object]") throw bridgeError("INVALID_PARAMS", "\u8BF7\u6C42\u53C2\u6570\u5FC5\u987B\u662F\u5BF9\u8C61");
  }
  function onlyKeys(params, allowed) {
    if (Object.keys(params).some((key) => !allowed.includes(key))) throw bridgeError("INVALID_PARAMS", "\u5305\u542B\u672A\u652F\u6301\u7684\u53C2\u6570");
  }
  function messageID() {
    const bytes = new Uint8Array(16);
    crypto.getRandomValues(bytes);
    return Array.from(bytes, (byte) => byte.toString(16).padStart(2, "0")).join("");
  }
  function scopedModel(adapter, requirements) {
    return {
      properties: requirements.filter((item) => item.kind === "property").map((item) => ({
        ...adapter.model.properties.find((field) => field.identifier === item.identifier),
        identifier: item.identifier,
        type: adapter.model.properties.find((field) => field.identifier === item.identifier)?.type || String(item.define?.type || ""),
        mode: adapter.model.properties.find((field) => field.identifier === item.identifier)?.mode || "r",
        available: adapter.model.properties.some(
          (field) => field.identifier === item.identifier && field.available !== false && matchesRequirement(item, field)
        ),
        panelAccess: item.access
      })),
      actions: requirements.filter((item) => item.kind === "action").map((item) => ({
        ...adapter.model.actions.find((field) => field.identifier === item.identifier),
        identifier: item.identifier,
        available: adapter.model.actions.some((field) => field.identifier === item.identifier && field.available !== false && matchesRequirement(item, field))
      }))
    };
  }
  function attachPanelBridge(frame, adapter, requirements, onError) {
    const completed = /* @__PURE__ */ new Map();
    const model = scopedModel(adapter, requirements);
    const sessionID = messageID();
    const subscriptions = /* @__PURE__ */ new Map();
    let disposed = false;
    let port;
    const send = (message) => {
      try {
        if (!disposed && port)
          port.postMessage(JSON.parse(JSON.stringify({ protocol: "ur-panel/v1", sessionID, ...message })));
      } catch {
        onError?.("\u9762\u677F\u54CD\u5E94\u65E0\u6CD5\u5E8F\u5217\u5316");
      }
    };
    const dispatch = async (method, params) => {
      requestObject(params);
      switch (method) {
        case "getAuth": {
          onlyKeys(params, []);
          if (!adapter.getAuth) throw bridgeError("UNSUPPORTED_CAPABILITY", "\u5F53\u524D\u5BBF\u4E3B\u672A\u63D0\u4F9B\u8BA4\u8BC1\u4FE1\u606F");
          return adapter.getAuth();
        }
        case "getContext": {
          onlyKeys(params, []);
          return {
            ...adapter.context,
            capabilities: {
              ...adapter.context.capabilities,
              settings: { available: !!adapter.openSettings }
            }
          };
        }
        case "getThingModel": {
          onlyKeys(params, []);
          return model;
        }
        case "invokeAction": {
          onlyKeys(params, ["identifier", "params"]);
          const action = model.actions.find((item) => item.identifier === params.identifier && item.available !== false);
          if (!action) throw bridgeError("FORBIDDEN", "\u884C\u4E3A\u672A\u58F0\u660E\u6216\u4E0D\u53EF\u7528");
          requestObject(params.params);
          const inputs = action.input || {};
          if (Object.keys(params.params).some((key) => !hasOwn(inputs, key)) || Object.keys(inputs).some((key) => !hasOwn(params.params, key)))
            throw bridgeError("INVALID_PARAMS", "\u884C\u4E3A\u8F93\u5165\u53C2\u6570\u4E0D\u5339\u914D");
          for (const [key, value] of Object.entries(params.params)) validateValue(inputs[key], value);
          return adapter.invokeAction(params.identifier, params.params);
        }
        case "openSettings": {
          onlyKeys(params, []);
          if (!adapter.openSettings) throw bridgeError("UNSUPPORTED_CAPABILITY", "\u5F53\u524D\u5BBF\u4E3B\u4E0D\u652F\u6301\u8BBE\u5907\u8BBE\u7F6E");
          await adapter.openSettings();
          return { opened: true };
        }
        case "readProperties": {
          onlyKeys(params, ["identifiers"]);
          const ids = params.identifiers ?? model.properties.filter((field) => field.available !== false).map((field) => field.identifier);
          if (!Array.isArray(ids) || ids.some((id) => typeof id !== "string" || !model.properties.some((field) => field.identifier === id)))
            throw bridgeError("INVALID_PARAMS", "\u672A\u77E5\u6216\u672A\u58F0\u660E\u7684\u5C5E\u6027");
          if (ids.some((id) => model.properties.find((field) => field.identifier === id)?.available === false))
            throw bridgeError("UNSUPPORTED_CAPABILITY", "\u5C5E\u6027\u4E0D\u53EF\u7528");
          const values = await adapter.readProperties(ids);
          requestObject(values);
          return Object.fromEntries(Object.entries(values).filter(([key]) => ids.includes(key)));
        }
        case "request": {
          validateApiRequest(params);
          if (!adapter.request) throw bridgeError("UNSUPPORTED_CAPABILITY", "\u5F53\u524D\u5BBF\u4E3B\u672A\u63D0\u4F9B\u5E73\u53F0\u63A5\u53E3\u8C03\u7528");
          return adapter.request(params);
        }
        case "subscribe": {
          onlyKeys(params, ["topics"]);
          if (!Array.isArray(params.topics) || params.topics.some((topic) => !["command", "connection", "property"].includes(topic)))
            throw bridgeError("INVALID_PARAMS", "\u8BA2\u9605\u4E3B\u9898\u65E0\u6548");
          const subscriptionID = messageID();
          subscriptions.set(
            subscriptionID,
            adapter.subscribe(params.topics, (event) => {
              if (!params.topics.includes(event.topic)) return;
              const data = event.topic === "property" ? Object.fromEntries(
                Object.entries(event.data).filter(([key]) => model.properties.some((field) => field.identifier === key && field.available !== false))
              ) : event.data;
              send({ type: "event", subscriptionID, topic: event.topic, data, timestamp: event.timestamp ?? Date.now() });
            })
          );
          return { subscriptionID };
        }
        case "unsubscribe": {
          onlyKeys(params, ["subscriptionID"]);
          subscriptions.get(params.subscriptionID)?.();
          subscriptions.delete(params.subscriptionID);
          return {};
        }
        case "writeProperties": {
          onlyKeys(params, ["values"]);
          requestObject(params.values);
          if (Object.keys(params.values).length === 0) throw bridgeError("INVALID_PARAMS", "\u5199\u5165\u503C\u4E0D\u80FD\u4E3A\u7A7A");
          for (const [id, value] of Object.entries(params.values)) {
            const field = model.properties.find((field2) => field2.identifier === id);
            if (!field || field.available === false || field.mode !== "rw" || field.panelAccess !== "write")
              throw bridgeError("FORBIDDEN", "\u5C5E\u6027\u672A\u58F0\u660E\u5199\u5165\u6216\u5F53\u524D\u4E0D\u53EF\u5199");
            if (field.type === "bool" && typeof value !== "boolean" || ["int", "timestamp"].includes(field.type) && !Number.isSafeInteger(value) || field.type === "float" && (typeof value !== "number" || !Number.isFinite(value)) || field.type === "string" && typeof value !== "string")
              throw bridgeError("INVALID_PARAMS", "\u5C5E\u6027\u6570\u636E\u7C7B\u578B\u65E0\u6548");
            const max = Number(field.max);
            if (typeof value === "string" && field.maxLen !== void 0 && value.length > Number(field.maxLen))
              throw bridgeError("INVALID_PARAMS", "\u5C5E\u6027\u5B57\u7B26\u4E32\u8FC7\u957F");
            const min = Number(field.min);
            if (typeof value === "number" && (Number.isFinite(min) && value < min || Number.isFinite(max) && value > max))
              throw bridgeError("INVALID_PARAMS", "\u5C5E\u6027\u8D85\u51FA\u5B9E\u9645\u8303\u56F4");
            if (field.type === "enum" && (!["number", "string"].includes(typeof value) || !hasOwn(field.mapping || field.values || {}, String(value))))
              throw bridgeError("INVALID_PARAMS", "\u679A\u4E3E\u503C\u65E0\u6548");
            if (!["bool", "enum", "float", "int", "string", "timestamp"].includes(field.type))
              throw bridgeError("UNSUPPORTED_CAPABILITY", "\u5F53\u524D\u672A\u5B9E\u73B0\u8BE5\u6570\u636E\u7C7B\u578B");
          }
          return adapter.writeProperties(params.values);
        }
        default: {
          throw bridgeError("UNSUPPORTED_CAPABILITY", "\u672A\u77E5\u7684\u9762\u677F\u65B9\u6CD5");
        }
      }
    };
    const onRequest = async (event) => {
      const message = event.data;
      if (disposed || !message || message.protocol !== "ur-panel/v1" || message.sessionID !== sessionID || message.type !== "request" || typeof message.id !== "string")
        return;
      let serialized;
      try {
        serialized = JSON.stringify(message);
      } catch {
        send({ type: "response", id: message.id, ok: false, error: { code: "INVALID_PARAMS", message: "\u8BF7\u6C42\u5FC5\u987B\u53EF\u5E8F\u5217\u5316" } });
        return;
      }
      if (serialized.length > 65536 || message.id.length > 128 || completed.size >= 1e3 && !completed.has(message.id)) {
        send({ type: "response", id: message.id, ok: false, error: { code: "INVALID_PARAMS", message: "\u8BF7\u6C42\u5927\u5C0F\u6216\u6570\u91CF\u8D85\u9650\uFF0C\u8BF7\u91CD\u65B0\u8FDB\u5165\u9762\u677F" } });
        return;
      }
      const fingerprint = JSON.stringify([message.method, message.params]);
      const prior = completed.get(message.id);
      if (prior && prior.fingerprint !== fingerprint) {
        send({ type: "response", id: message.id, ok: false, error: { code: "INVALID_PARAMS", message: "\u8BF7\u6C42 ID \u5DF2\u7528\u4E8E\u5176\u4ED6\u53C2\u6570" } });
        return;
      }
      const response = prior?.response ?? Promise.resolve().then(() => dispatch(message.method, message.params ?? {})).then(
        (result) => ({ type: "response", id: message.id, ok: true, result }),
        (error) => ({
          type: "response",
          id: message.id,
          ok: false,
          error: { code: error.code || "INVALID_PARAMS", message: error.message || "\u9762\u677F\u64CD\u4F5C\u5931\u8D25" }
        })
      );
      if (!prior) completed.set(message.id, { fingerprint, response });
      send(await response);
    };
    const ready = (event) => {
      if (disposed || port || event.source !== frame.contentWindow || event.data?.protocol !== "ur-panel/v1" || event.data.type !== "ready") return;
      const channel = new MessageChannel();
      port = channel.port1;
      port.addEventListener("message", onRequest);
      port.start();
      frame.contentWindow?.postMessage({ protocol: "ur-panel/v1", type: "init", sessionID, context: adapter.context }, "*", [channel.port2]);
      clearTimeout(timer);
    };
    const timer = setTimeout(() => {
      if (!port && !disposed) onError?.("\u9762\u677F SDK \u521D\u59CB\u5316\u8D85\u65F6");
    }, 5e3);
    window.addEventListener("message", ready);
    return () => {
      if (disposed) return;
      send({ type: "close" });
      disposed = true;
      clearTimeout(timer);
      window.removeEventListener("message", ready);
      for (const cleanup of subscriptions.values()) cleanup();
      subscriptions.clear();
      completed.clear();
      port?.close();
      adapter.dispose?.();
    };
  }
  function validateValue(define, value) {
    const type = define?.type;
    if (type === "bool" && typeof value === "boolean" || type === "string" && typeof value === "string") return;
    if (type === "enum" && ["number", "string"].includes(typeof value) && hasOwn(define.mapping || define.values || {}, String(value))) return;
    if (["float", "int", "timestamp"].includes(type) && typeof value === "number" && Number.isFinite(value) && (type === "float" || Number.isSafeInteger(value)) && (define.min === null || define.min === void 0 || value >= Number(define.min)) && (define.max === null || define.max === void 0 || value <= Number(define.max)))
      return;
    throw bridgeError("INVALID_PARAMS", "\u884C\u4E3A\u8F93\u5165\u503C\u4E0D\u7B26\u5408\u5B9E\u9645\u7269\u6A21\u578B");
  }
  function matchesRequirement(requirement, field) {
    if (requirement.kind === "action") return canonical(requirement.input || {}) === canonical(field.input || {});
    return ["type", "unit", "mapping", "values"].every(
      (key) => requirement.define?.[key] === void 0 || canonical(requirement.define[key]) === canonical(field[key])
    );
  }
  function hasOwn(value, key) {
    return Object.prototype.hasOwnProperty.call(value, key);
  }
  function canonical(value) {
    if (Array.isArray(value)) return `[${value.map((item) => canonical(item)).join(",")}]`;
    if (value && typeof value === "object")
      return `{${Object.keys(value).sort().map((key) => `${JSON.stringify(key)}:${canonical(value[key])}`).join(",")}}`;
    return JSON.stringify(value) ?? "undefined";
  }

  // apps/web/packages/saas/src/device-panel/mock-settings.ts
  function createMockSettings(context) {
    let dialog;
    const dispose = () => {
      const current = dialog;
      dialog = void 0;
      current?.close();
      current?.remove();
    };
    const open = () => {
      if (dialog?.open) {
        dialog.focus();
        return;
      }
      dialog = document.createElement("dialog");
      dialog.setAttribute("aria-label", "\u6A21\u62DF\u8BBE\u5907\u8BBE\u7F6E");
      dialog.style.cssText = "box-sizing:border-box;width:min(440px,calc(100vw - 32px));max-height:85vh;overflow:auto;padding:24px;border:1px solid color-mix(in srgb,currentColor 16%,transparent);border-radius:16px;background:Canvas;color:CanvasText;font:14px/1.6 system-ui,sans-serif;box-shadow:0 16px 60px color-mix(in srgb,CanvasText 15%,transparent)";
      const title = document.createElement("h2");
      title.textContent = "\u8BBE\u5907\u8BBE\u7F6E";
      title.style.cssText = "margin:0 0 8px;font-size:20px";
      const note = document.createElement("p");
      note.textContent = "\u6A21\u62DF\u8BBE\u5907 \xB7 \u4EC5\u9884\u89C8\u8BBE\u7F6E\u5165\u53E3\uFF0C\u4E0D\u8BFB\u53D6\u6216\u4FEE\u6539\u771F\u5B9E\u8BBE\u5907";
      note.style.cssText = "margin:0 0 20px;opacity:.65";
      const list = document.createElement("dl");
      for (const [label, value] of [
        ["\u8BBE\u5907\u540D\u79F0", context.displayName || context.deviceName],
        ["\u8BBE\u5907 ID", context.deviceID],
        ["\u4EA7\u54C1 ID", context.productID]
      ]) {
        const row = document.createElement("div");
        row.style.cssText = "display:flex;justify-content:space-between;gap:24px;padding:16px 0;border-top:1px solid color-mix(in srgb,currentColor 10%,transparent)";
        const name = document.createElement("dt");
        const text = document.createElement("dd");
        name.textContent = String(label);
        text.textContent = String(value || "\u6682\u65E0\u4FE1\u606F");
        text.style.cssText = "margin:0;text-align:right;overflow-wrap:anywhere";
        row.append(name, text);
        list.append(row);
      }
      const help = document.createElement("p");
      help.textContent = "\u771F\u5B9E\u73AF\u5883\u7531\u5BBF\u4E3B\u63D0\u4F9B\u8BBE\u5907\u4FE1\u606F\u3001\u540D\u79F0\u548C\u5176\u4ED6\u53EF\u7528\u8BBE\u7F6E\uFF1B\u6743\u9650\u4E0E\u5F53\u524D\u8BBE\u5907\u4FDD\u6301\u4E00\u81F4\u3002";
      help.style.cssText = "opacity:.65;margin:16px 0 24px";
      const close = document.createElement("button");
      close.type = "button";
      close.textContent = "\u8FD4\u56DE\u9762\u677F";
      close.style.cssText = "width:100%;min-height:44px;border:0;border-radius:8px;background:Highlight;color:HighlightText;font:inherit;cursor:pointer";
      close.addEventListener("click", dispose);
      const current = dialog;
      current.addEventListener(
        "close",
        () => {
          current.remove();
          if (dialog === current) dialog = void 0;
        },
        { once: true }
      );
      dialog.append(title, note, list, help, close);
      document.body.append(dialog);
      dialog.showModal();
      close.focus();
    };
    return { open, dispose };
  }

  // apps/web/packages/saas/src/device-panel/mock.ts
  function createMockAdapter(requirements, options = {}) {
    const listeners = /* @__PURE__ */ new Set();
    const values = {};
    const properties = requirements.filter((item) => item.kind === "property").map((item) => ({
      ...item.define,
      identifier: item.identifier,
      type: String(item.define?.type || "string"),
      mode: item.access === "write" ? "rw" : "r",
      available: true
    }));
    for (const field of properties) {
      let value = "";
      if (field.type === "bool") value = false;
      else if (["float", "int", "timestamp"].includes(field.type)) value = Number(field.min || 0);
      else if (field.type === "enum") value = Object.keys(field.mapping || field.values || {})[0] ?? "";
      values[field.identifier] = value;
    }
    const emit = (event) => listeners.forEach((callback) => callback(event));
    const snapshot = (ids) => Object.fromEntries(ids.map((id) => [id, { value: values[id], timestamp: Date.now(), receivedTime: Date.now(), quality: "good", source: "mock" }]));
    const context = {
      mode: "mock",
      deviceID: "mock-device",
      productID: "mock-product",
      deviceName: "\u6A21\u62DF\u8BBE\u5907",
      capabilities: {
        settings: { available: true },
        cloud: { available: true },
        api: { available: true, source: "mock" },
        ble: { available: false, reason: "\u6A21\u62DF\u5BBF\u4E3B\u672A\u8FDE\u63A5\u84DD\u7259" }
      }
    };
    const settings = createMockSettings(context);
    const adapter = {
      context,
      openSettings: settings.open,
      model: {
        properties,
        actions: requirements.filter((item) => item.kind === "action").map((item) => ({ identifier: item.identifier, input: item.input || {}, available: true }))
      },
      async getAuth() {
        return { baseURL: "", headers: {} };
      },
      async request(params) {
        validateApiRequest(params);
        const handler = options.apiHandlers?.[params.path];
        if (!handler) throw bridgeError("UNSUPPORTED_CAPABILITY", "\u6A21\u62DF\u63A5\u53E3\u5C1A\u672A\u914D\u7F6E\uFF0C\u4E0D\u4F1A\u8BBF\u95EE\u771F\u5B9E\u8BBE\u5907");
        return handler(params.data || {});
      },
      async readProperties(ids) {
        return snapshot(ids);
      },
      async writeProperties(next) {
        Object.assign(values, next);
        const commandID = messageID();
        emit({ topic: "property", data: snapshot(Object.keys(next)) });
        setTimeout(() => emit({ topic: "command", data: { commandID, status: "succeeded", source: "mock" } }), 50);
        return { commandID, status: "accepted", source: "mock" };
      },
      async invokeAction(identifier, params) {
        const commandID = messageID();
        setTimeout(() => emit({ topic: "command", data: { commandID, identifier, status: "succeeded", output: params, source: "mock" } }), 50);
        return { commandID, status: "accepted", source: "mock" };
      },
      subscribe(topics, callback) {
        const scoped = (event) => {
          if (topics.includes(event.topic)) callback(event);
        };
        listeners.add(scoped);
        return () => listeners.delete(scoped);
      },
      dispose() {
        settings.dispose();
        listeners.clear();
      }
    };
    return { adapter, emit, values };
  }

  // apps/web/packages/saas/src/device-panel/dev-host.ts
  document.querySelectorAll("[data-size]").forEach((button) => {
    button.addEventListener("click", () => {
      const preview = document.querySelector("#preview");
      if (preview) preview.dataset.size = button.dataset.size;
      document.querySelectorAll("[data-size]").forEach((item) => item.setAttribute("aria-pressed", String(item === button)));
    });
  });
  document.querySelector("#reload")?.addEventListener("click", () => location.reload());
  async function start() {
    const frame = document.querySelector("#panel");
    const error = document.querySelector("#error");
    if (!frame || !error) return;
    try {
      const response = await fetch("/__ur_panel/manifest");
      if (!response.ok) throw new Error("\u5DE5\u7A0B\u6821\u9A8C\u5931\u8D25\uFF0C\u8BF7\u8FD0\u884C ur things panel validate");
      const manifest = await response.json();
      const { adapter } = createMockAdapter(manifest.requirements);
      const cleanup = attachPanelBridge(frame, adapter, manifest.requirements, (text) => {
        error.textContent = text;
      });
      window.addEventListener("pagehide", cleanup, { once: true });
      frame.src = "/panel/index.html";
    } catch (error_) {
      error.textContent = error_.message || "\u6A21\u62DF\u5BBF\u4E3B\u542F\u52A8\u5931\u8D25";
    }
  }
  void start();
})();
