/* 文件说明：真实建筑计量范围导航、进线总功率、区域功率色阶及按需历史展示。 */
(function(){
  'use strict';
  const config=window.BuildingConfig,data=window.BuildingData; // 空间展示与真实遥测分离。
  const byId=id=>document.getElementById(id),format=n=>n===null?'—':n.toFixed(2);
  const escape=s=>String(s).replace(/[&<>"']/g,c=>({'&':'&amp;','<':'&lt;','>':'&gt;','"':'&quot;',"'":'&#39;'}[c])); // 平台别名作为文本，不能执行 HTML。
  const state={mode:'exterior',expanded:false,spacing:4,floor:null,zone:null,scope:null};
  let scene=null; // WebGL 不可用时，真实计量列表仍可操作。
  /** current 确定唯一计量范围；地下总箱关系未核实时仅分别展示，不作整层汇总。 */
  function current(){
    const floor=config.floors.find(f=>f.id===state.floor),zone=floor?.zones.find(z=>z.id===state.zone);
    const group=config.groups.find(g=>g.id===(state.scope||zone?.groupId||floor?.sharedGroup));
    const summary=group?data.aggregate([group]):floor?data.aggregate(floor.zones):data.total();
    if(floor&&!group&&!floor.canSum){summary.power=null;summary.noSum=true;summary.history=[];}
    return {floor,zone,group,summary};
  }
  /** load 当前选择变化时才请求对应历史，实时刷新不触发重复查询。 */
  function load(){const {summary}=current();data.select(summary.noSum?[]:summary.ids);}
  /** select 保留已验证的单击切层与镜头规则；跨层共用回路仍为同一个数据范围。 */
  function select(floor,zone,scope=null){
    const changed=state.floor!==floor||!state.expanded;Object.assign(state,{floor,zone,scope});
    if(floor){state.mode='energy';state.expanded=true;}
    scene?.setState(state,changed&&!!floor);refresh();load();
  }
  /** trend 用实际时间轴绘制断点和负功率，缺少历史不伪造零线。 */
  function trend(summary){
    const points=summary.history||[],valid=points.filter(p=>p.v!==null);
    if(!valid.length){byId('trend').innerHTML='<div class="empty">'+(summary.noSum?'各回路分别计量':summary.historyLoading?'正在读取历史':summary.historyError?'历史读取失败':'暂无有效历史')+'<small>近 24 小时 · 每 30 分钟对齐</small></div>';return;}
    const min=Math.min(0,...valid.map(p=>p.v)),max=Math.max(1,...valid.map(p=>p.v))*1.1,range=max-min;
    const start=points[0].t,end=points.at(-1).t; // 所有成员使用同一时间点，断点处重新开始路径。
    let path='',connected=false;for(const p of points){if(p.v===null){connected=false;continue;}const x=32+(p.t-start)/(end-start)*230,y=100-(p.v-min)/range*82;path+=(connected?'L':'M')+x.toFixed(2)+','+y.toFixed(2);connected=true;}
    const time=t=>new Date(t).toLocaleTimeString('zh-CN',{hour:'2-digit',minute:'2-digit',hour12:false});
    const grid=[0,.5,1].map(t=>'<line x1="32" x2="262" y1="'+(100-t*82)+'" y2="'+(100-t*82)+'" stroke="#29404e" stroke-dasharray="2 4"/><text x="26" y="'+(103-t*82)+'" text-anchor="end">'+(min+t*range).toFixed(1)+'</text>').join('');
    byId('trend').innerHTML='<svg viewBox="0 0 280 124" preserveAspectRatio="none"><g fill="#7894a4" font-size="8" font-family="monospace">'+grid+'<text x="32" y="120">'+time(start)+'</text><text x="262" y="120" text-anchor="end">'+time(end)+'</text></g><path d="'+path+'" stroke="#76d9c8" stroke-width="1.8" fill="none"/></svg>';
    byId('trend').setAttribute('aria-label','近24小时真实功率趋势，缺失数据保留断点');
  }
  /** refresh 数据更新只刷新读数和色阶，不移动镜头或重建楼层。 */
  function refresh(){
    const {floor,zone,group,summary}=current(),total=data.total(),coverage=data.summarize(config.meters.map(m=>m.id));
    byId('total-power').textContent=format(total.power);byId('coverage').textContent=coverage.known+' / '+coverage.total;
    byId('detail-title').textContent=group?group.name:floor?floor.id+' · '+floor.name:'整栋建筑';
    byId('detail-breadcrumb').textContent='BUILDING / '+(group?group.id:floor?.id||'INCOMING');
    byId('detail-subtitle').textContent=group?.shared?'跨层合并计量 · 不作分层分摊':group?'供电回路 '+group.id:summary.noSum?'总分关系待核对 · 各回路分别显示':floor?'已关联回路功率':config.incoming.length?'已配置进线总功率':'进线计量待配置';
    byId('detail-power').textContent=summary.noSum?'分别计量':format(summary.power);
    byId('detail-power').style.fontSize=summary.noSum?'28px':'';
    byId('detail-coverage').textContent=data.mode!=='real'?(data.mode==='invalid'?'同一设备的计量配置冲突':data.mode==='unconfigured'?'功率属性或单位未配置':data.mode==='unbound'?'未配置电表绑定':'等待平台数据连接'):summary.stale?'含 '+summary.stale+' 块表数据已过期':summary.known<summary.total?'有效 '+summary.known+'/'+summary.total+' 块表'+(summary.error?' · 部分请求失败':' · 数据未齐'):'已接入 '+summary.total+' 块表';
    byId('detail-capacity').textContent=summary.total+' 块';
    byId('detail-share').textContent=summary.power===null||total.power===null||total.power<=0?'—':(summary.power/total.power*100).toFixed(1)+'%';
    byId('load-ratio').textContent=summary.known+'/'+summary.total;
    byId('load-fill').style.width=(summary.total?summary.known/summary.total*100:0)+'%';byId('load-fill').style.background='#76d9c8';
    byId('load-note').textContent=summary.ts?'最早数据时间 '+new Date(summary.ts).toLocaleString('zh-CN',{hour12:false}):'等待带时间戳的功率记录';
    byId('scene-title').textContent=group?group.name:floor?floor.id+' · 空间用能':state.expanded?'分层探索':state.mode==='energy'?'功耗分布':'建筑全景';
    byId('scene-number').textContent=state.expanded?'02 / EXPLODED FLOORS':'01 / BUILDING OVERVIEW';
    byId('scene-subtitle').textContent=group?.shared?'跨层计量范围 · 同一回路只计一次':state.mode==='exterior'?'白色曲面 · 竖向幕墙 · 公共裙房':'点击楼层与回路，查看对应功率';
    for(const mode of ['exterior','energy']){byId('mode-'+mode).classList.toggle('active',state.mode===mode);byId('mode-'+mode).setAttribute('aria-pressed',String(state.mode===mode));}
    byId('explode').classList.toggle('active',state.expanded);byId('explode').setAttribute('aria-pressed',String(state.expanded));byId('explode').innerHTML=state.expanded?'合拢建筑 <span>↕</span>':'分层展开 <span>↕</span>';
    byId('spacing').disabled=!state.expanded;byId('spacing').value=state.spacing;byId('spacing-value').textContent=state.spacing.toFixed(1);byId('all-floors').classList.toggle('selected',!floor&&!group);
    for(const f of config.floors){const button=document.querySelector('[data-floor="'+f.id+'"]'),s=data.aggregate(f.zones);button.classList.toggle('selected',state.floor===f.id);button.setAttribute('aria-pressed',String(state.floor===f.id));button.querySelector('.floor-number').textContent=f.shared?'合并计量':!f.canSum?'分别计量':format(s.power);}
    const bucket=byId('extra-category').value; // 屋顶、室外和未定位回路不强行落到楼层。
    const children=floor?floor.zones.map(z=>({id:z.id,name:z.name,summary:data.aggregate([z]),floor:floor.id,zone:z.id})):bucket?config.groups.filter(g=>!g.floors.length&&g.category===bucket).map(g=>({id:g.id,name:g.name+' · '+g.id,summary:data.aggregate([g]),scope:g.id})):config.floors.filter((f,i,all)=>!f.sharedGroup||!all.slice(0,i).some(prev=>prev.sharedGroup===f.sharedGroup)).map(f=>({id:f.id,name:f.shared?config.groups.find(g=>g.id===f.sharedGroup).name:f.id+' · '+f.name,summary:f.canSum?data.aggregate(f.zones):{power:null},floor:f.id}));
    byId('children-title').textContent=floor?'本层计量回路':bucket?'独立回路读数':'计量范围';
    byId('children-list').replaceChildren(...children.map(c=>{const button=document.createElement('button');button.className='child-row';button.dataset.child=c.id;button.innerHTML='<span>'+escape(c.name)+'</span><b>'+format(c.summary.power)+'</b>';button.onclick=()=>select(c.floor||null,c.zone||null,c.scope||null);return button;}));
    byId('evidence').textContent=group?'电表 '+group.meters.join('、')+(group.shared?' · 共用计量范围，不代表任一单层功率':''):floor?'点击外露楼板切层，点击区域查看该回路。':'仅汇总已核验的进线范围；空间轮廓为可替换示例。';
    const max=data.maximum();byId('power-legend').textContent=max>0?'0 / '+[.25,.5,.75,1].map(t=>(max*t).toFixed(2)).join(' / ')+' kW':config.groups.some(g=>data.aggregate([g]).power===0)?'0 kW':'暂无有效功率';
    trend(summary);scene?.refreshData();
  }
  config.floors.slice().reverse().forEach(f=>{const b=document.createElement('button');b.className='floor-row';b.dataset.floor=f.id;b.innerHTML='<span class="floor-id">'+f.id+'</span><span class="floor-name">'+escape(f.name)+'</span><span class="floor-number">—</span>';b.onclick=()=>{byId('extra-category').value='';select(f.id,null);};byId('floor-list').appendChild(b);});
  byId('all-floors').onclick=()=>{byId('extra-category').value='';select(null,null);};
  byId('extra-category').onchange=()=>{Object.assign(state,{floor:null,zone:null,scope:null});scene?.setState(state);refresh();load();};
  /** reset 统一恢复全景，不重置已经加载的真实遥测。 */
  function reset(){Object.assign(state,{mode:'exterior',expanded:false,spacing:4,floor:null,zone:null,scope:null});byId('extra-category').value='';scene?.setState(state);refresh();load();}
  byId('reset').onclick=reset;byId('mode-exterior').onclick=reset;
  byId('mode-energy').onclick=()=>{Object.assign(state,{mode:'energy',expanded:true});scene?.setState(state);refresh();};
  byId('explode').onclick=()=>{Object.assign(state,{expanded:!state.expanded,floor:null,zone:null,scope:null});state.mode=state.expanded?'energy':'exterior';scene?.setState(state);refresh();load();};
  byId('spacing').oninput=e=>{state.spacing=Number(e.target.value);scene?.setState(state);refresh();};
  document.title=config.name+' · 建筑能耗';document.querySelector('h1').firstChild.textContent=config.name;
  refresh();try{scene=window.createBuildingScene(byId('viewport'),select);}catch(e){console.error('建筑渲染初始化失败',e);byId('load-text').textContent='请启用浏览器硬件加速';}
  data.onChange(refresh);
  // 让主体首帧先提交，再启动绑定和数据请求，避免同步初值挤占首屏绘制。
  requestAnimationFrame(()=>setTimeout(()=>data.start(),0));
  window.BuildingApp={scene,inspect:()=>({...state,summary:current().summary,scene:scene?.inspect(),dataMode:data.mode,bound:data.bound}),get total(){return data.total();}}; // 只读验收入口不暴露凭证。
})();
