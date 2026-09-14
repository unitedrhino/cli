/* 文件说明：部署前编辑的建筑模板配置；只有空间示例，不包含设备绑定或现场计量事实。 */
window.BuildingSettings = {
  name: '建筑能耗中心', // 页面名称，可按部署场景修改。
  telemetry: { identifier: '', factor: null, staleMs: 75 * 60000, refreshMs: 60000 }, // 核对物模型后设置；factor 将源值换算成 kW。
  incoming: [], // 经确认的进线本地锚点 ID；空数组不计算整楼功率，禁止与馈电混加。
  summableFloors: [], // 已确认无总分表重叠的楼层 ID；默认各范围分别计量。
  floorGroups: {}, // 用实际计量范围 ID 覆盖楼层区域引用；多层引用同组即共享计量，不做比例分摊。
  groups: [], // 可添加 {id,name,meters:[本地锚点ID],category}；只表达经核实的计量关系。
  meters: [], // 可添加 {id,name,telemetry:{identifier,factor,staleMs}}，产品/设备来自宿主绑定表。
};
