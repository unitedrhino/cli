/* 文件说明：模板脱离项目仓库后的页面与桥接回归；测试夹具只在浏览器测试上下文注入。 */
const {chromium}=require('playwright'),fs=require('node:fs'),path=require('node:path'),{pathToFileURL}=require('node:url'),assert=require('node:assert/strict');
const out=path.resolve(process.argv.includes('--out')?process.argv[process.argv.indexOf('--out')+1]:path.join(__dirname,'.temp/verification')); // 截图输出不参与发布。
/** fixture 注入无凭证的测试宿主，并在加载配置时显式建立测试计量关系。 */
function fixture(){
  let settings,init,data; // 仅用于测试的配置、初始化和增量回调。
  Object.defineProperty(window,'BuildingSettings',{get:()=>settings,set:v=>{settings=v;v.incoming=['main-a','main-b'];v.telemetry.identifier='activePower';v.telemetry.factor=.001;v.floorGroups={'5F':['level-5-a'],'6F':['level-5-a']};}});
  const sdk={onInit:f=>init=f,onData:f=>data=f,subscribe:()=>{},reportAnchors:()=>{},ready:()=>init(BuildingConfig.meters.map((m,i)=>({path:'meter/'+m.id,productID:'test-product',deviceName:'device-'+i}))),callApi:async(p,b)=>{
    if(window.__fail)throw Error('测试断网');const value=b.deviceName==='device-0'?5000:b.deviceName==='device-1'?8000:0;
    return {list:[{dataID:'activePower',value,timestamp:Date.now()}],total:1};
  }};
  Object.defineProperty(window,'UrSceneSDK',{get:()=>sdk,set:()=>{}});window.__update=v=>data(v);
}
/** ready 等待真实构建完成与镜头收敛。 */
async function ready(p){await p.waitForFunction(()=>window.BuildingApp?.scene?.metrics.ready&&!BuildingApp.scene.inspect().animating);}
(async()=>{fs.mkdirSync(out,{recursive:true});const browser=await chromium.launch({headless:true,args:['--no-sandbox','--enable-unsafe-swiftshader','--use-angle=swiftshader']});try{
  const p=await browser.newPage({viewport:{width:1920,height:1080}}),errors=[];p.on('pageerror',e=>errors.push(e.message));
  await p.goto(pathToFileURL(path.join(__dirname,'index.html')).href);await ready(p);
  assert.equal(await p.locator('#total-power').innerText(),'—');assert.equal(await p.evaluate(()=>BuildingData.bound),0);
  await p.screenshot({path:path.join(out,'unbound.png')});
  await p.click('[data-floor="3F"]');await ready(p);await p.locator('.zone-label:visible').first().click();assert.ok(await p.evaluate(()=>BuildingApp.inspect().zone));assert.equal(await p.locator('#detail-power').innerText(),'—');
  await p.click('[data-floor="7F"]');await ready(p);assert.equal(await p.evaluate(()=>BuildingApp.inspect().floor),'7F');
  for(const width of [1239,720]){await p.setViewportSize({width,height:1080});await ready(p);assert.equal(await p.evaluate(()=>document.documentElement.scrollWidth>innerWidth),false);}
  await p.close();
  const live=await browser.newPage({viewport:{width:1920,height:1080}});live.on('pageerror',e=>errors.push(e.message));await live.addInitScript(fixture);await live.goto(pathToFileURL(path.join(__dirname,'index.html')).href);await ready(live);
  await live.waitForFunction(()=>BuildingData.total().power===13);await live.click('[data-floor="5F"]');await ready(live);assert.match(await live.locator('#detail-subtitle').innerText(),/跨层合并/);assert.equal(await live.locator('#detail-power').innerText(),'0.00');
  await live.click('[data-floor="6F"]');await ready(live);assert.equal(await live.locator('#detail-power').innerText(),'0.00');
  await live.evaluate(()=>__update({productID:'test-product',deviceName:'device-12',identifier:'activePower',value:-1000,ts:Date.now()+100}));await live.waitForFunction(()=>BuildingApp.inspect().summary.power===-1);
  await live.screenshot({path:path.join(out,'shared-range-test.png')});
  const metrics=await live.evaluate(()=>BuildingApp.scene.inspect());assert.deepEqual(errors,[]);fs.writeFileSync(path.join(out,'report.json'),JSON.stringify({metrics,errors,cases:['无绑定无读数','区域与跨层单击','窄屏','宿主属性与单位','共享计量','负功率']},null,2));console.log('页面通过，截图与指标：'+out);
}finally{await browser.close();}})().catch(e=>{console.error(e);process.exitCode=1;});
