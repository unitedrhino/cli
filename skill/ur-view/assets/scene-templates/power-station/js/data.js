/* 文件说明：通用配电站示例布局与真实接入配置；本地编号仅用于稳定锚点，不代表设备或供电范围。 */
(function () {
  'use strict';
  /** 接入配置：先核验物模型，再填 identifier 与换算到展示单位的 multiplier；空标识符不查询。 */
  window.STATION_DATA = {
    staleMs: 75 * 60000,
    pollMs: 60000,
    historyStepMs: 30 * 60000,
    fields: {
      ua: { identifier: '', multiplier: 1, unit: 'V' },
      ub: { identifier: '', multiplier: 1, unit: 'V' },
      uc: { identifier: '', multiplier: 1, unit: 'V' },
      ia: { identifier: '', multiplier: 1, unit: 'A' },
      ib: { identifier: '', multiplier: 1, unit: 'A' },
      ic: { identifier: '', multiplier: 1, unit: 'A' },
      p: { identifier: '', multiplier: 1, unit: 'kW' },
      q: { identifier: '', multiplier: 1, unit: 'kvar' },
      pf: { identifier: '', multiplier: 1, unit: '' },
      kwh: { identifier: '', multiplier: 1, unit: 'kWh' },
      on: { identifier: '', type: 'boolean', trueValue: 1, falseValue: 0 }
    },
    /** 只有确认计量关系后才能启用；成员必须是进线锚点，不能混入馈电或分表。 */
    summary: { verified: false, meterIds: [] }
  };
  /** 页面标题和设备结构均为可替换的通用示例。 */
  window.STATION = { name: '配电站三维运维', voltage: '', meta: {} };
  /** 高压柜仅提供外观，不预设电压等级、电源或现场设备身份。 */
  window.HV_CABINETS = Array.from({ length: 6 }, (_, i) => ({
    id: 'HV-' + String(i + 1).padStart(2, '0'), name: '高压柜 ' + (i + 1),
    kind: ['hv-in', 'hv-pt', 'hv-out'][i % 3], feeder: '', breaker: ''
  }));
  /** 变压器不附推测容量、铭牌额定值和实际供电关系。 */
  window.TRANSFORMERS = ['A', 'B'].map(id => ({ id: 'TR-' + id, name: '变压器 ' + id, model: '干式变压器', capacity: '', ratio: '', vector: '' }));
  /** 两列柜体保留不同抽屉密度供尺度评审；这些数量与空间位置没有计量语义。 */
  window.LV_SECTIONS = [[13, 12, 11, 9, 10, 11], [10, 9, 10, 8, 11, 12]].map((counts, i) => {
    const side = ['A', 'B'][i]; // 柜列内部编号。
    return {
      id: 'S' + (i + 1), name: '低压柜列 ' + side, side,
      cabinets: [
        { id: side + '-IN', name: '进线柜 ' + side, kind: 'incoming', meterId: side + '-IN-M', meterType: '多功能电力仪表' },
        { id: side + '-COMP', name: '无功补偿柜 ' + side, kind: 'svg' },
        ...counts.map((count, j) => {
          const id = side + '-CAB-' + String(j + 1).padStart(2, '0'); // 柜体稳定编号。
          return { id, name: '馈电柜 ' + side + (j + 1), kind: 'feeder', prefix: id,
            drawers: Array.from({ length: count }, (_, k) => ({ no: k + 1, id: id + '-M' + String(k + 1).padStart(2, '0'), circuit: '仪表 ' + side + (j + 1) + '-' + (k + 1) })) };
        })
      ]
    };
  });
  /** 母联只表示模型构件，不参与默认数据汇总。 */
  window.BUS_CABINET = { id: 'BUS-LINK', name: '母联柜', kind: 'bus', meterType: '' };
  /** 从示例布局生成锚点清单；返回条目仅含本地名称，不包含产品或设备身份。 */
  window.buildMeterRegistry = function () {
    return window.LV_SECTIONS.flatMap(sec => sec.cabinets.flatMap(cab => {
      const common = { cabinet: cab.id, cabinetName: cab.name, kind: cab.kind, meterType: '多功能电力仪表', section: sec.name }; // 所属结构信息。
      if (cab.kind === 'incoming') return [{ ...common, id: cab.meterId, name: '进线仪表 ' + sec.side }];
      return (cab.drawers || []).map(drawer => ({ ...common, id: drawer.id, name: drawer.circuit }));
    }));
  };
})();
