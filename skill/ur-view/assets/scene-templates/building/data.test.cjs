/* 文件说明：宿主绑定、单位、唯一设备、乱序和换绑回归；测试值仅在 VM 注入，不进入运行包。 */
const fs=require('node:fs'),vm=require('node:vm'),path=require('node:path'),assert=require('node:assert/strict');
/** tick 等待请求队列排空，避免用固定接口延迟掩盖竞态。 */
const tick=async()=>{for(let i=0;i<20;i++)await new Promise(setImmediate);};
(async()=>{
  const requests=[],handlers={},now=Date.now(),bucket=Math.floor(now/1800000)*1800000; // 受控物模型夹具与消息回调。
  let fail=false,release=null; // 模拟失败和换绑前的在途响应，只属于测试。
  const sdk={onInit:f=>handlers.init=f,onData:f=>handlers.data=f,subscribe:(_,f)=>handlers.prop=f,reportAnchors:()=>{},ready:()=>{},callApi:async(p,b)=>{
    requests.push({p,b});if(fail)throw Error('fixture unavailable');
    if(b.deviceName==='delayed')return new Promise(r=>release=()=>r({list:[{dataID:'watts',value:999000,timestamp:Date.now()}]}));
    const value=b.productID==='product-a'?3000:7000;
    return {list:[{dataID:'watts',value,timestamp:p.includes('property-log')?bucket:now}],total:1};
  }};
  const c={setTimeout:(f,ms)=>ms===20000?0:setTimeout(f,ms),clearTimeout,setInterval:()=>0,clearInterval:()=>{},Date,Map,Set,UrSceneSDK:sdk};c.window=c;vm.createContext(c);
  const load=n=>vm.runInContext(fs.readFileSync(path.join(__dirname,'js',n+'.js'),'utf8'),c);
  load('settings');c.BuildingSettings.telemetry.identifier='watts';c.BuildingSettings.telemetry.factor=.001;load('config');load('data');
  const d=c.BuildingData,cfg=c.BuildingConfig;
  assert.equal(d.total().power,null);assert.equal(cfg.incoming.length,0);assert.ok(cfg.floors.every(f=>!f.canSum&&!f.shared));
  d.start();d.init([]);await tick();assert.equal(requests.length,0,'未绑定不调用设备接口');assert.ok([...d.states.values()].every(s=>s.power===null));
  cfg.incoming.push('main-a','main-b');
  const list=[{path:'meter/main-a',productID:'product-a',deviceName:'same-name'},{path:'meter/main-b',productID:'product-b',deviceName:'same-name'},{path:'meter/level-1-a',productID:'product-a',deviceName:'same-name'}];
  d.init(list);await tick();assert.equal(d.total().power,10,'跨产品同名分别计量且 W 换算 kW');
  assert.equal(d.summarize(['main-a','level-1-a']).power,3,'同设备两个锚点不重复累计');
  assert.equal(requests.length,2,'重复锚点仅查询一次快照');
  d.init(list);await tick();assert.equal(requests.length,2,'重复握手不重复加载');
  handlers.data({productID:'product-a',deviceName:'same-name',identifier:'watts',value:0,ts:now+10});assert.equal(d.summarize(['main-a']).power,0);
  handlers.data({productID:'unknown',deviceName:'same-name',identifier:'watts',value:8000,ts:now+11});assert.equal(d.summarize(['main-a']).power,0);
  handlers.prop({device:{productID:'product-a',deviceName:'same-name'},params:{watts:-2000},timestamp:now+20});assert.equal(d.summarize(['main-a']).power,-2);
  assert.equal(d.color(d.loadRatio(d.summarize(['main-a']))),'#69a9ed');
  assert.equal(d.ingest('main-a',9000,now),false);assert.equal(d.ingest('main-a','',now+30),false);assert.equal(d.ingest('main-a',null,now+30),false);
  assert.equal(d.summarize(['main-a'],now+cfg.staleMs+100).power,null);
  d.select(['main-a','level-1-a','main-b']);await tick();assert.equal(requests.filter(r=>r.p.includes('property-log')).length,2);
  assert.equal(d.total().history.at(-1).v,10);assert.equal(d.total().history.at(-2).v,null);
  assert.ok(requests.every(r=>r.b.productID&&r.b.deviceName),'所有API均限定绑定设备');
  fail=true;d.refresh();await tick();assert.ok(d.states.get('main-a').error);fail=false;d.refresh();await tick();assert.equal(d.states.get('main-a').error,false);
  d.init([{path:'meter/main-a',productID:'product-a',deviceName:'delayed'}]);await tick();assert.equal(d.states.get('main-a').power,null);
  d.init([{path:'meter/main-a',productID:'product-b',deviceName:'replacement'}]);await tick();assert.equal(d.states.get('main-a').power,7);release();await tick();assert.equal(d.states.get('main-a').power,7,'换绑后旧请求不能倒灌');
  d.init([...list,{path:'meter/main-a',productID:'other-product',deviceName:'other-device'}]);await tick();assert.equal(d.mode,'invalid');assert.equal(d.total().power,null,'歧义锚点必须清空数据');
  d.init([]);await tick();assert.equal(d.total().power,null);assert.equal(d.bound,0);
  console.log('通过：空绑定、配置进线、跨产品同名、别名去重、单位、零负过期、乱序、按需历史、失败恢复、换绑隔离');
})().catch(e=>{console.error(e);process.exitCode=1;});
