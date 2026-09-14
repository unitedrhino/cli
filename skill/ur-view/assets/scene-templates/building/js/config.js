/* 文件说明：将独立配置转为楼层与计量范围；中性锚点仅作占位，真实设备由宿主绑定。 */
(function(root){
  'use strict';
  const settings=root.BuildingSettings; // 唯一部署配置入口。
  const floorIds=['B1','1F','2F','3F','4F','5F','6F','7F','8F']; // 外形示例的展示层，非现场层数结论。
  const meters=[{id:'main-a',name:'进线位置 A',kind:'incoming'},{id:'main-b',name:'进线位置 B',kind:'incoming'}]; // 保留进线绑定入口，不默认加入总功率。
  const groups=[]; // 仅为示例区块提供一对一锚点，不预置设备身份。
  floorIds.forEach((floor,i)=>['a','b'].forEach((part,j)=>{
    const id='level-'+i+'-'+part;meters.push({id,name:floor+' · 区域 '+(j+1),kind:'feeder'});
    groups.push({id,name:floor+' · 区域 '+(j+1),meters:[id],category:'unassigned'});
  }));
  ['roof','outdoor','unassigned'].forEach(category=>{const id='extra-'+category;meters.push({id,name:'独立范围 · '+category,kind:'feeder'});groups.push({id,name:category==='roof'?'屋面范围':category==='outdoor'?'室外范围':'待定位范围',meters:[id],category});});
  settings.meters.forEach(m=>{const i=meters.findIndex(v=>v.id===m.id);if(i<0)meters.push(m);else meters[i]={...meters[i],...m};});
  settings.groups.forEach(g=>{const i=groups.findIndex(v=>v.id===g.id);if(i<0)groups.push(g);else groups[i]=g;});
  const floors=floorIds.map((id,index)=>{
    const refs=settings.floorGroups[id]||['level-'+index+'-a','level-'+index+'-b']; // 无配置时仅显示中性占位区域。
    const width=index<=2?43:30-Math.max(0,index-6)*1.8,depth=index<=2?27:18,centerZ=index<=2?0:-3.5;
    const columns=refs.length===1?1:refs.length>4?3:2,rows=Math.max(1,Math.ceil(refs.length/columns));
    refs.forEach(ref=>{if(!groups.some(g=>g.id===ref))throw Error('未知计量范围 '+ref);});
    return {id,name:'楼层空间',index,y:(index-1)*3.8,width,depth,centerZ,canSum:settings.summableFloors.includes(id),
      zones:refs.map((ref,i)=>({id:id+'/'+ref,groupId:ref,name:groups.find(g=>g.id===ref).name,floorId:id,
        x:-width/2+(i%columns+.5)*width/columns,z:centerZ-depth/2+(Math.floor(i/columns)+.5)*depth/rows,
        width:width/columns-1.1,depth:depth/rows-1.1,reference:ref})),
    };
  });
  groups.forEach(g=>{g.floors=floors.filter(f=>f.zones.some(z=>z.groupId===g.id)).map(f=>f.id);g.shared=g.floors.length>1;});
  floors.forEach(f=>{f.sharedGroup=f.zones.length===1&&groups.find(g=>g.id===f.zones[0].groupId).shared?f.zones[0].groupId:null;f.shared=!!f.sharedGroup;});
  for(const id of [...settings.incoming,...groups.flatMap(g=>g.meters)])if(!meters.some(m=>m.id===id))throw Error('未知锚点 '+id);
  settings.incoming.forEach(id=>{if(meters.find(m=>m.id===id).kind!=='incoming')throw Error('进线汇总不能包含馈电锚点 '+id);});
  root.BuildingConfig={name:settings.name,meters,incoming:settings.incoming,groups,floors,telemetry:settings.telemetry,staleMs:settings.telemetry.staleMs};
})(typeof window==='undefined'?globalThis:window);
