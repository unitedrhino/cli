/* 文件说明：宿主真实数据桥、设备去重与历史汇总；不含模拟读数或自动回退。
 * 锚点与设备分离，只有完整 productID/deviceName 身份匹配才更新；换绑使旧在途结果失效。
 */
(function(root){
  'use strict';
  const config=root.BuildingConfig,bindings=new Map(),records=new Map(),listeners=new Set(); // 锚点映射、唯一设备状态与界面订阅。
  const states=new Map(),historyCache=new Map(),pending=new Set(),tasks=[]; // 对外锚点状态、设备缓存和有界请求队列。
  let sdk=null,active=0,closed=false,timer=null,notifyTimer=null,mode='waiting',generation=0,signature=''; // 数据桥生命周期。
  const selected=new Set(config.incoming); // 仅请求当前范围历史，整楼初始范围由配置指定。
  /** empty 每次生成独立空记录，避免将未连接解释成零。 */
  function empty(){return {power:null,ts:0,points:[],error:false};}
  config.meters.forEach(m=>states.set(m.id,empty()));
  /** key 产品与设备组合构成唯一身份，跨产品同名设备不混淆。 */
  function key(b){return JSON.stringify([String(b.productID),String(b.deviceName)]);}
  /** metric 每个锚点可覆盖默认物模型标识、倍率和有效期。 */
  function metric(id){return {...config.telemetry,...config.meters.find(m=>m.id===id)?.telemetry};}
  /** timestamp 兼容秒/毫秒，缺失时间不伪装当前记录。 */
  function timestamp(v){const n=Number(v);return Number.isFinite(n)&&n>0?(n<1e11?n*1000:n):0;}
  /** number 空值、布尔和非数字不能作为零功率。 */
  function number(v){return (typeof v==='number'||typeof v==='string'&&v.trim()!=='')&&Number.isFinite(Number(v))?Number(v):null;}
  /** unique 依据绑定设备去重；未绑定锚点仍是独立缺失成员。 */
  function unique(ids){const seen=new Set();return ids.filter(id=>{const k=bindings.has(id)?key(bindings.get(id)):'slot:'+id;if(seen.has(k))return false;seen.add(k);return true;});}
  /** emit 同批遥测合并刷新，避免挤占三维主线程。 */
  function emit(){if(!notifyTimer&&!closed)notifyTimer=setTimeout(()=>{notifyTimer=null;listeners.forEach(fn=>fn());},80);}
  /** ingest 接收原始物模型值并换算 kW，严格拒绝倒序与无效时间。 */
  function ingest(id,value,time){const s=states.get(id),m=metric(id),raw=number(value),ts=timestamp(time),power=raw===null?null:raw*m.factor;
    if(!bindings.has(id)||!s||power===null||!Number.isFinite(power)||!ts||ts>Date.now()+60000||ts<=s.ts)return false;
    s.power=power;s.ts=ts;s.error=false;s.points=[...s.points.filter(p=>p.t>=Date.now()-86400000-m.staleMs),{t:ts,v:power}].sort((a,b)=>a.t-b.t);emit();return true;
  }
  /** idsFor 图形范围映射为本地锚点，绝不按面积分配功率。 */
  function idsFor(zones){return unique(zones.flatMap(z=>z.meters||config.groups.find(g=>g.id===(z.groupId||z.id))?.meters||[]));}
  /** summarize 同一设备只计一次；缺成员时总值为空，历史断点不补零。 */
  function summarize(input,now=Date.now()){
    const ids=unique(input),valid=ids.filter(id=>{const s=states.get(id);return s&&s.power!==null&&now-s.ts<=metric(id).staleMs;});
    const partial=valid.length?valid.reduce((sum,id)=>sum+states.get(id).power,0):null,step=1800000,end=Math.floor(now/step)*step;
    const history=Array.from({length:49},(_,i)=>{const t=end-(48-i)*step;let sum=0;for(const id of ids){const p=states.get(id)?.points.filter(p=>p.t<=t).at(-1);if(!p||t-p.t>metric(id).staleMs)return {t,v:null};sum+=p.v;}return {t,v:ids.length?sum:null};});
    return {power:ids.length&&valid.length===ids.length?partial:null,partial,known:valid.length,total:ids.length,ids,history,
      stale:ids.filter(id=>states.get(id)?.ts&&now-states.get(id).ts>metric(id).staleMs).length,ts:ids.length?Math.min(...ids.map(id=>states.get(id)?.ts||0)):0,
      historyLoading:ids.some(id=>historyCache.get(bindings.has(id)?key(bindings.get(id)):id)?.pending),
      historyError:ids.some(id=>historyCache.get(bindings.has(id)?key(bindings.get(id)):id)?.error),error:ids.some(id=>states.get(id)?.error)};
  }
  /** aggregate 汇总配置中的唯一计量范围。 */
  function aggregate(zones){return summarize(idsFor(zones));}
  /** maximum 颜色基准仅考虑当前有效馈电正功率，不表达额定容量或告警。 */
  function maximum(){return Math.max(0,...config.groups.map(g=>{let sum=0;for(const id of unique(g.meters)){const s=states.get(id);if(!s||s.power===null||Date.now()-s.ts>metric(id).staleMs)return 0;sum+=s.power;}return sum;}));}
  /** loadRatio 兼容渲染接口，比例仅为当前最大功率的相对值。 */
  function loadRatio(s){return s.power===null?null:s.power<0?-1:maximum()>0?s.power/maximum():0;}
  /** color 缺失灰色、反向蓝色、正向四档。 */
  function color(r){return r===null?'#728294':r<0?'#69a9ed':r<.25?'#4bcaa2':r<.5?'#e3c866':r<.75?'#ed965a':'#ef6873';}
  const values={}; // 区域 getter 始终读最新遥测。
  config.floors.forEach(f=>f.zones.forEach(z=>Object.defineProperty(values,z.id,{get:()=>aggregate([z])})));
  /** enqueue 限制快照和历史合计最多三路在途请求。 */
  function enqueue(fn){return new Promise((resolve,reject)=>{tasks.push({fn,resolve,reject});drain();});}
  /** drain 请求执行前再核验换绑代际，旧队列不会访问新设备。 */
  function drain(){while(active<3&&tasks.length&&!closed){const task=tasks.shift();active++;Promise.resolve().then(task.fn).then(task.resolve,task.reject).finally(()=>{active--;drain();});}}
  /** request 只通过宿主代调，服务端错误转为明确状态。 */
  async function request(path,body){const r=await sdk.callApi(path,body,{timeout:20000});if(r?.code&&r.code!==200)throw Error(r.msg||'请求失败');return r?.data??r;}
  /** fetchLatest 按绑定设备读取指定属性，不扫描产品内其他设备。 */
  async function fetchLatest(id){const b=bindings.get(id);if(!sdk||!b||closed)return;const k=key(b),g=generation,taskKey=g+':'+k,m=metric(id);
    if(pending.has(taskKey))return;pending.add(taskKey);
    try{await enqueue(async()=>{if(g!==generation)return;const r=await request('/api/v1/things/device/msg/property-latest/get-list',{productID:b.productID,deviceName:b.deviceName});if(g!==generation)return;
      for(const p of r?.list||[])if(p.dataID===m.identifier)ingest(id,p.value,p.timestamp);states.get(id).error=false;});}
    catch(_){if(g===generation)states.get(id).error=true;}finally{pending.delete(taskKey);emit();}
  }
  /** fetchHistory 当前选择按设备缓存一分钟，历史统一换算后保留时间轴。 */
  async function fetchHistory(id){const b=bindings.get(id);if(!sdk||!b||closed)return;const k=key(b),cache=historyCache.get(k),g=generation,m=metric(id);
    if(cache?.pending||cache&&!cache.error&&Date.now()-cache.at<60000)return;
    historyCache.set(k,{pending:true,at:Date.now(),error:false});emit();
    try{const points=[],now=Date.now();for(let page=1;page<=10;page++){
      const r=await enqueue(()=>g===generation?request('/api/v1/things/device/msg/property-log/get-list',{productID:b.productID,deviceName:b.deviceName,dataID:m.identifier,timeStart:String(now-86400000-m.staleMs),timeEnd:String(now),order:2,page:{page,size:200}}):null);
      if(g!==generation)return;
      for(const p of r?.list||[]){const raw=number(p.value),t=timestamp(p.timestamp);if(raw!==null&&t&&t<=now)points.push({t,v:raw*m.factor});}
      if(!r?.list?.length||page*200>=Number(r.total||r.list.length))break;if(page===10)throw Error('历史查询超限');
    }
    const s=states.get(id),map=new Map(points.map(p=>[p.t,p]));s.points.forEach(p=>map.set(p.t,p));s.points=[...map.values()].filter(p=>p.t>=now-86400000-m.staleMs).sort((a,b)=>a.t-b.t);historyCache.set(k,{pending:false,at:now,error:false});
    }catch(_){if(g===generation)historyCache.set(k,{pending:false,at:Date.now(),error:true});}emit();
  }
  /** select 导航时只加载当前范围历史，缺快照则补读。 */
  function select(ids){selected.clear();unique(ids).forEach(id=>{selected.add(id);if(!states.get(id)?.ts)fetchLatest(id);fetchHistory(id);});}
  /** invalidate 拒绝歧义配置并立即清空旧设备状态，在途结果随代际失效。 */
  function invalidate(){generation++;signature='';bindings.clear();records.clear();historyCache.clear();states.forEach((_,id)=>states.set(id,empty()));mode='invalid';emit();}
  /** init 宿主绑定为唯一设备来源；重复握手保留数据，换绑立即清除旧值。 */
  function init(list){
    const next=new Map(),schemas=new Map();let missingMetric=false; // 同设备绑定到多个锚点时，属性/倍率必须一致，否则拒绝歧义配置。
    for(const b of Array.isArray(list)?list:[]){if(typeof b?.path!=='string'||!b.path.startsWith('meter/'))continue;const id=b.path.slice(6);if(!states.has(id)||b.productID==null||!String(b.productID)||!b.deviceName)continue;
      const m=metric(id);if(!m.identifier||!Number.isFinite(m.factor)||m.factor<=0||!(m.staleMs>0)){missingMetric=true;continue;}
      const mapped={productID:String(b.productID),deviceName:String(b.deviceName)},k=key(mapped),schema=JSON.stringify(m);
      if(schemas.has(k)&&schemas.get(k)!==schema){invalidate();return;}if(next.has(id)&&key(next.get(id))!==k){invalidate();return;}schemas.set(k,schema);next.set(id,mapped);
    }
    const sig=JSON.stringify([...next.entries()].sort(([a],[b])=>a.localeCompare(b)));mode=next.size?'real':missingMetric?'unconfigured':'unbound';emit();if(sig===signature)return;
    signature=sig;generation++;const old=new Map(bindings);bindings.clear();next.forEach((b,id)=>bindings.set(id,b));historyCache.clear();
    const oldRecords=new Map(records);records.clear();states.forEach((_,id)=>{const b=bindings.get(id);if(!b){states.set(id,empty());return;}const k=key(b),prior=old.get(id);const s=records.get(k)||(prior&&key(prior)===k?oldRecords.get(k):null)||empty();records.set(k,s);states.set(id,s);});
    unique(config.incoming).forEach(fetchLatest);select([...selected]);unique([...bindings.keys()]).filter(id=>!config.incoming.includes(id)).forEach(fetchLatest);
  }
  /** receive 只认产品/设备组合匹配的消息，将属性标识交给每个锚点配置。 */
  function receive(productID,deviceName,identifier,value,ts){for(const [id,b] of bindings)if(String(productID)===b.productID&&deviceName===b.deviceName&&identifier===metric(id).identifier)ingest(id,value,ts);}
  /** refresh 定期补读缺失/过期快照并重试当前范围，持续断网仍不阻塞模型。 */
  function refresh(){emit();unique([...config.incoming,...selected]).forEach(id=>{fetchLatest(id);fetchHistory(id);});unique([...bindings.keys()]).forEach(id=>{const s=states.get(id);if(!s.ts||Date.now()-s.ts>metric(id).staleMs)fetchLatest(id);});}
  /** start SDK 只持有消息桥，场景不读取或保存登录凭证。 */
  function start(){sdk=root.UrSceneSDK;if(!sdk){mode='unbound';emit();return;}
    sdk.onInit(init);sdk.onData(p=>receive(p.productID,p.deviceName,p.identifier,p.value,p.ts));
    sdk.subscribe('prop',p=>{for(const [id,b] of bindings){if(String(p?.device?.productID)!==b.productID||p.device.deviceName!==b.deviceName)continue;const field=metric(id).identifier;ingest(id,p.params?.[field]??(p.identifier===field?p.param:undefined),p.timestamp);}});
    setTimeout(()=>{if(mode==='waiting'){mode='unbound';emit();}},20000);sdk.reportAnchors(config.meters.map(m=>({path:'meter/'+m.id,name:m.name})));sdk.ready();
    timer=setInterval(refresh,Math.max(1000,config.telemetry.refreshMs));root.addEventListener?.('pagehide',()=>{closed=true;clearInterval(timer);clearTimeout(notifyTimer);});
  }
  root.BuildingData={values,aggregate,summarize,loadRatio,color,maximum,select,start,ingest,init,states,historyCache,refresh,
    onChange:fn=>listeners.add(fn),get mode(){return mode;},get bound(){return bindings.size;},total:()=>summarize(config.incoming)};
})(typeof window==='undefined'?globalThis:window);
