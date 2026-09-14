/* 文件说明：仅在测试进程中注入虚拟绑定与读数，验证真实数据桥边界；测试文件不进入发布包。 */
const fs = require('node:fs');
const path = require('node:path');
const vm = require('node:vm');
const assert = require('node:assert/strict');
/** 创建隔离浏览器上下文，不访问网络。 */
function setup(callApi = async () => ({ code: 200, data: { list: [] } })) {
  const window = { parent: {}, UrSceneSDK: { callApi } }; // 仅测试宿主。
  const context = vm.createContext({ window, console, setInterval: () => 1, clearInterval: () => {}, Date, Map, Set }); // 定时器不延长测试生命周期。
  for (const file of ['data.js', 'urscene-provider.js']) vm.runInContext(fs.readFileSync(path.join(__dirname, 'js', file), 'utf8'), context);
  const meters = window.buildMeterRegistry(); // 示例布局锚点。
  window.STATION_DATA.fields.p.identifier = 'active_power';
  window.STATION_DATA.fields.p.multiplier = 0.001;
  window.STATION_DATA.fields.on.identifier = 'switch_state';
  window.STATION_DATA.fields.kwh.identifier = 'energy';
  const provider = new window.PowerProviders.UrSceneProvider(meters); // 数据桥待测实例。
  return { window, meters, provider };
}
/** 等待全部请求队列结束。 */
async function settle(provider) { while (provider.pending.size) await Promise.all([...provider.pending.values()]); }
/** 创建只用于测试的锚点绑定，两个不同产品允许使用相同设备名。 */
function binding(id, productID = 'test-product-a', deviceName = 'test-device') { return { path: 'meter/' + id, productID, deviceName }; }
/** 推送测试数据并触发可见状态刷新。 */
function send(provider, power, time, productID = 'test-product-a', deviceName = 'test-device') {
  provider._handleData({ productID, deviceName, identifier: 'active_power', value: power, ts: time });
}
/** 执行边界验收。 */
async function run() {
  const calls = [], source = setup(async (url, body) => { calls.push({ url, body }); return { code: 200, data: { list: [] } }; }); // 网络请求审计。
  const { provider, meters } = source, now = Date.now(); // 固定本轮时间。
  const unconfiguredCalls = [], unconfigured = setup(async (url, body) => { unconfiguredCalls.push({ url, body }); return { list: [] }; }); // 无字段映射不得调用接口。
  Object.values(unconfigured.provider.config.fields).forEach(field => { field.identifier = ''; });
  unconfigured.provider._handleInit([binding('A-IN-M')]); await settle(unconfigured.provider);
  assert.equal(unconfiguredCalls.length, 0); assert.equal(unconfigured.provider.get('A-IN-M').status, 'unconfigured');
  const unwrapped = setup(async url => ({ list: url.includes('property-log') ? [{ timestamp: Date.now() - 3600000, value: 1000 }] : [{ dataID: 'active_power', timestamp: Date.now(), value: 3000 }], total: 1 })); // SDK 已解包响应同样可用。
  unwrapped.provider._handleInit([binding('A-IN-M')]); await settle(unwrapped.provider);
  assert.equal(unwrapped.provider.get('A-IN-M').p, 3);
  await unwrapped.provider.fetchTrend('A-IN-M', 'p'); assert.equal(unwrapped.provider.get('A-IN-M').trendStatus.p, 'ready');
  const switchState = source.window.PowerProviders.UrSceneProvider.switchState; // 树节点状态不得由电压或功率推断。
  assert.equal(switchState({ status: 'unbound', on: null }), 'unknown');
  assert.equal(switchState({ status: 'ready', on: null, p: 1 }), 'unknown');
  assert.equal(switchState({ status: 'ready', on: true }), 'on');
  assert.equal(switchState({ status: 'ready', on: false }), 'off');
  assert.equal(switchState({ status: 'stale', on: true }), 'unknown');
  assert.equal(meters.length, 128);
  assert.equal(provider.get('A-IN-M').p, null);
  assert.equal(provider.get('A-IN-M').on, null);
  assert.equal(provider.summary('p'), null);
  provider._handleInit([]); await settle(provider);
  assert.equal(provider.mode, 'unbound'); assert.equal(calls.length, 0);
  provider._handleInit([binding('A-IN-M'), binding('B-IN-M'), binding('A-CAB-01-M01', 'test-product-b')]);
  await settle(provider);
  assert.equal(calls.length, 2, '同设备的两个锚点只查询一次');
  assert(calls.every(call => call.body.productID && call.body.deviceName && !call.url.includes('property-log')));
  send(provider, '', now); send(provider, false, now); send(provider, 'bad', now); send(provider, 30, null);
  assert.equal(provider.get('A-IN-M').p, null);
  send(provider, 2500, now);
  assert.equal(provider.get('A-IN-M').p, 2.5);
  assert.equal(provider.get('B-IN-M').p, 2.5);
  assert.equal(provider.get('A-CAB-01-M01').p, null, '不同产品同名设备不串值');
  send(provider, 0, now + 1); send(provider, 9000, now);
  assert.equal(provider.get('A-IN-M').p, 0, '零值有效且旧时间戳不得覆盖');
  send(provider, -2000, now + 2);
  assert.equal(provider.get('A-IN-M').p, -2);
  provider.config.summary = { verified: true, meterIds: ['A-IN-M', 'B-IN-M'] };
  assert.equal(provider.summary('p'), -2, '汇总按唯一设备去重');
  provider.config.summary.meterIds.push('A-CAB-01-M01');
  assert.equal(provider.summary('p'), null, '馈电不能混入进线汇总');
  provider.config.staleMs = 1;
  await new Promise(resolve => setTimeout(resolve, 5)); provider._notify();
  assert.equal(provider.get('A-IN-M').p, null); assert.equal(provider.get('A-IN-M').status, 'stale');
  provider._handleInit([binding('A-IN-M', 'test-product-b')]); await settle(provider);
  assert.equal(provider.get('B-IN-M').status, 'unbound');
  send(provider, 9999, now + 4); assert.equal(provider.get('A-IN-M').p, null, '旧设备推送不得污染新绑定');
  provider._handleInit([binding('A-IN-M'), binding('A-IN-M', 'test-product-b')]);
  assert.equal(provider.mode, 'unbound', '同锚点歧义绑定不猜测');
  const requests = [], history = setup(async (url, body) => { requests.push({ url, body }); return { code: 200, data: { list: url.includes('property-log') ? [{ timestamp: now - 2 * 3600000, value: 1000 }, { timestamp: now, value: 0 }] : [] } }; }); // 历史时间断点夹具。
  history.provider._handleInit([binding('A-IN-M')]); await settle(history.provider);
  await history.provider.fetchTrend('A-IN-M', 'p');
  assert.equal(requests.filter(request => request.url.includes('property-log')).length, 1);
  await history.provider.fetchTrend('A-IN-M', 'p');
  assert.equal(requests.filter(request => request.url.includes('property-log')).length, 1, '一分钟历史缓存');
  assert(history.provider.get('A-IN-M').trends.p.some(point => point.v === null), '过期采样保留历史断点');
  assert(history.provider.get('A-IN-M').trends.p.some(point => point.v === 1), '历史采用相同单位换算');
  let resolveOld; // 暂停旧绑定快照，验证在途隔离。
  const race = setup((url, body) => body.productID === 'test-product-a' ? new Promise(resolve => { resolveOld = resolve; }) : Promise.resolve({ code: 200, data: { list: [] } }));
  race.provider._handleInit([binding('A-IN-M')]); await Promise.resolve();
  race.provider._handleInit([binding('A-IN-M', 'test-product-b')]);
  resolveOld({ code: 200, data: { list: [{ dataID: 'active_power', timestamp: Date.now(), value: 8888 }] } });
  await settle(race.provider);
  assert.equal(race.provider.get('A-IN-M').p, null, '旧绑定在途快照不得覆盖新绑定');
  let fail = true; // 失败后立即重试历史。
  const recovery = setup(async url => {
    if (url.includes('property-log') && fail) return { code: 500 };
    return { code: 200, data: { list: [] } };
  });
  recovery.provider._handleInit([binding('A-IN-M')]); await settle(recovery.provider);
  await recovery.provider.fetchTrend('A-IN-M', 'p'); assert.equal(recovery.provider.get('A-IN-M').trendStatus.p, 'error');
  fail = false; await recovery.provider.fetchTrend('A-IN-M', 'p'); assert.equal(recovery.provider.get('A-IN-M').trendStatus.p, 'empty');
  const stopped = setup(); stopped.provider.stop(); // 退出后不再消费队列。
  console.log('配电站数据桥：128 锚点、无绑定、严格数值、去重、单位、时序、汇总、过期、历史断点、在途绑定隔离、错误恢复全部通过。');
}
run().catch(error => { console.error(error); process.exitCode = 1; });
