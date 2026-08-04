#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
gen-view-components.py — 盘点 GoView 大屏全部组件，生成 CLI 组件数据 JSON 与 skill 组件清单 Markdown。

用途：
  1. 遍历 goview 组件目录，从各组件 index.ts 正则提取 ConfigType 元数据
     （key/chartKey/conKey/title/category/categoryName/package/chartFrame/image）；
     仅统计经分组 index.ts 注册、可被 packagesList 到达的组件（未注册的遗留目录自动排除）。
     package/category/chartFrame 输出枚举**值**（如 Charts/Bars/echarts），
     与大屏画布存档 JSON 中实际存储的字符串一致。
  2. 结合前端 IoT 支持矩阵（queryTypeSupport.ts 的 isQueryTypeSupported 逻辑）
     与 IOT_SINGLE_VALUE_SELF_RENDER_CHARTS 单值自渲染集合，计算每个组件的
     iotQueryTypes 与 singleValueSelfRender。
  3. 产出：
     - components.json：供 ur view screen validate/describe 命令 go:embed 使用（schema 固定）；
     - components.md：ur-view skill 的组件清单参考文档。

用法：
  python3 scripts/gen-view-components.py \
      [--src <goview components 目录>] \
      [--json-out <components.json 输出路径>] \
      [--md-out <components.md 输出路径>] \
      [--snapshot-date <快照日期 YYYY-MM-DD>]

说明：
  - 幂等：同一源码快照重复运行产出完全一致。
  - IoT 支持矩阵硬编码自以下来源（快照 2026-08-03），源码变更后需同步更新：
      apps/web/packages/bigscreen/src/goview/views/chart/ContentConfigurations/components/ChartData/components/ChartDataIoTDevice/queryTypeSupport.ts
      apps/web/packages/bigscreen/src/goview/hooks/useChartDataFetch.hook.ts（IOT_SINGLE_VALUE_SELF_RENDER_CHARTS）
  - 只读取前端源码，不做任何修改。
"""

import argparse
import json
import os
import re
import sys

# ── 8 大组件包注册顺序（与 packages/index.ts 的 packagesList 一致） ──
PACKAGE_ORDER = [
    'Charts',
    'Informations',
    'Tables',
    'Decorates',
    'Icons',
    'Photos',
    'Presets',
    'Interact',
]

PACKAGE_CN = {
    'Charts': '图表',
    'Informations': '信息',
    'Tables': '表格',
    'Decorates': '装饰',
    'Icons': '图标',
    'Photos': '图片',
    'Presets': '预置',
    'Interact': '交互',
}

# ── IoT queryType 支持矩阵：镜像 queryTypeSupport.ts 的组件集合（2026-08-03 快照） ──
LINE_CHART_KEYS = {'LineCommon', 'LineGradients', 'LineGradientSingle', 'LineLinearSingle'}
PIE_CHART_KEYS = {'PieCircle', 'PieCommon'}
BAR_CHART_KEYS = {'BarCommon', 'BarCrossrange', 'BarLine'}
CAPSULE_CHART_KEYS = {'CapsuleChart', 'CapsuleCommon'}
TEXT_CHART_KEYS = {'TextCommon', 'FlipperNumber'}
DEVICE_INFO_CHART_KEYS = {'TablesBasic', 'TableList'}
OPEN_LIST_CHART_KEYS = {'TablesBasic', 'TableScrollBoard'}
RANK_LIST_CHART_KEYS = {'TableList'}
RADAR_CHART_KEYS = {'Radar'}
PROGRESS_CHART_KEYS = {'Process', 'WaterPolo'}

# ── 单值自渲染集合：镜像 useChartDataFetch.hook.ts 的 IOT_SINGLE_VALUE_SELF_RENDER_CHARTS ──
IOT_SINGLE_VALUE_SELF_RENDER_CHARTS = {'VPieCircle', 'VProcess', 'VWaterPolo', 'VDial'}

# ── 经典数据面板限定 静态/AJAX 的组件（ChartData/index.vue filteredOptions），不走通用 IoT 面板 ──
IOT_CLASSIC_ONLY_KEYS = {'BorderDeviceStatus', 'BorderRunningTimer'}

# ── 分组容器（chartKey=group），不参与数据绑定 ──
CONTAINER_KEYS = {'BorderContainer'}

# ── 适用场景一句话（按 key 精确匹配，缺省回退到包级描述） ──
SCENARIOS = {
    'BarCommon': '分类数值对比柱状图',
    'BarCrossrange': '横向条形排名对比',
    'BarLine': '柱线混合双轴对比',
    'CapsuleChart': '多设备同属性最新值胶囊柱对比',
    'LineCommon': '属性历史时序趋势',
    'LineGradients': '渐变面积趋势图',
    'LineGradientSingle': '单系列渐变趋势',
    'LineLinearSingle': '单系列线性趋势',
    'PieCommon': '组成占比饼图',
    'PieCircle': '环形占比 / 单值百分比',
    'Process': '单值进度条',
    'WaterPolo': '单值水位球',
    'ScatterCommon': '双数值散点分布',
    'ScatterLogarithmicRegression': '对数回归散点',
    'MapBase': '地理区域分布图',
    'MapAmap': '高德地图点位展示',
    'Radar': '多维指标雷达对比',
    'Funnel': '流程转化漏斗',
    'Heatmap': '密度 / 矩阵热力图',
    'TreeMap': '层级占比矩形树图',
    'Graph': '关系拓扑图',
    'Sankey': '流量去向桑基图',
    'Dial': '单值仪表盘',
    'TextCommon': '单值文本 / 单项信息展示',
    'TextGradient': '渐变标题文字',
    'TextBarrage': '滚动弹幕信息',
    'InputsDate': '日期筛选控件',
    'InputsSelect': '下拉筛选控件',
    'InputsTab': '页签切换控件',
    'InputsPagination': '分页控件',
    'InputsInput': '输入框控件',
    'Image': '图片展示',
    'ImageCarousel': '图片轮播',
    'Video': '视频播放',
    'Iframe': '内嵌第三方网页',
    'WordCloud': '关键词词云',
    'TableList': '设备信息滚动排行列表',
    'TableScrollBoard': '滚动数据表格',
    'TablesBasic': '基础表格 / 多设备字段对比',
    'Number': '大数字单值展示',
    'FlipperNumber': '翻牌器单值展示',
    'TimeCommon': '当前时间展示',
    'CountDown': '倒计时',
    'Clock': '时钟展示',
    'FullScreen': '全屏切换按钮',
    'PipelineH': '横向管道装饰',
    'PipelineV': '纵向管道装饰',
    'CirclePoint': '圆点装饰',
    'ThreeEarth01': '三维地球展示',
    'Icon': '矢量图标（图标库统一入口）',
    'BorderRunningTimer': '边框容器 + 设备累计运行时间',
    'BorderMetricFlipper': '边框容器 + 指标翻牌器',
    'BorderDeviceStatus': '边框容器 + 设备状态统计',
    'BorderDualChart': '边框容器 + 双设备对比图表',
    'BorderBarChart': '边框容器 + 通用柱状图',
    'BorderContainer': '空白边框分组容器',
    'DeviceInteractButton': '设备控制交互按钮',
}

PACKAGE_SCENARIO_DEFAULT = {
    'Charts': '通用图表组件',
    'Informations': '信息展示组件',
    'Tables': '表格组件',
    'Decorates': '纯装饰组件',
    'Icons': '图标组件',
    'Photos': '图片组件',
    'Presets': '业务复合预置组件',
    'Interact': '交互组件',
}


def read_text(path):
    with open(path, 'r', encoding='utf-8') as f:
        return f.read()


def parse_enum_maps(index_type_path):
    """解析包目录 index.type.ts 中的枚举，返回 {枚举名: {成员名: 字符串值}}。"""
    maps = {}
    if not os.path.isfile(index_type_path):
        return maps
    text = read_text(index_type_path)
    for m in re.finditer(r'export\s+enum\s+(\w+)\s*\{(.*?)\}', text, re.S):
        enum_name, body = m.group(1), m.group(2)
        members = {}
        for em in re.finditer(r"(\w+)\s*=\s*'([^']*)'", body):
            members[em.group(1)] = em.group(2)
        maps[enum_name] = members
    return maps


def extract_literal(text, field):
    """从 ConfigType 对象文本中提取 `field: 'value'` 字面量。"""
    m = re.search(r"\b%s:\s*'([^']*)'" % field, text)
    return m.group(1) if m else ''


def extract_enum_member(text, field, enum_names):
    """提取 `field: <枚举名>.<成员>` 的成员名。"""
    for enum_name in enum_names:
        m = re.search(r'\b%s:\s*%s\.(\w+)' % (field, enum_name), text)
        if m:
            return m.group(1)
    return ''


def resolve_index(src_dir, base_file, spec):
    """把 import 相对路径解析为组件/分组目录的 index.ts 绝对路径；无法解析返回 None。"""
    if not spec.startswith('.'):
        return None
    if spec.endswith('/config'):
        return None
    base_dir = os.path.dirname(base_file)
    target = os.path.normpath(os.path.join(base_dir, spec))
    if os.path.isfile(target + '.ts'):
        # 直接指向某个 ts 文件（非目录），仅在是 index.ts 时有效
        if target.endswith('/index'):
            return target + '.ts'
        return None
    if os.path.isdir(target):
        candidate = os.path.join(target, 'index.ts')
        if os.path.isfile(candidate):
            return candidate
    return None


def find_registered_components(src):
    """从 8 个包根 index.ts 出发做 BFS，按注册顺序收集组件 index.ts（按 key 去重）。"""
    components = []  # [{package, dir, index_ts}]
    seen_keys = set()
    for package in PACKAGE_ORDER:
        root = os.path.join(src, package, 'index.ts')
        if not os.path.isfile(root):
            continue
        queue = [root]
        visited = set()
        while queue:
            current = queue.pop(0)
            if current in visited:
                continue
            visited.add(current)
            text = read_text(current)
            for m in re.finditer(r"from\s*['\"]([^'\"]+)['\"]", text):
                resolved = resolve_index(src, current, m.group(1))
                if not resolved or resolved in visited:
                    continue
                resolved_text = read_text(resolved)
                if re.search(r"\bkey:\s*'", resolved_text):
                    # 组件文件：提取 key 去重
                    key = extract_literal(resolved_text, 'key')
                    if key and key not in seen_keys:
                        seen_keys.add(key)
                        components.append({
                            'package': package,
                            'dir': os.path.dirname(resolved),
                            'index_ts': resolved,
                            'text': resolved_text,
                        })
                else:
                    queue.append(resolved)
    return components


def has_dataset(config_ts_path):
    """判断 config.ts 的 option 中是否声明了 dataset 属性（剔除注释后匹配）。"""
    if not os.path.isfile(config_ts_path):
        return False
    text = read_text(config_ts_path)
    text = re.sub(r'/\*.*?\*/', '', text, flags=re.S)
    text = re.sub(r'(?m)//.*$', '', text)
    return re.search(r'(?m)^\s*dataset\s*:', text) is not None


def compute_iot_query_types(key, chart_key, chart_frame, dataset_present):
    """按 queryTypeSupport.ts 的 isQueryTypeSupported 逻辑计算支持的 IoT queryType 列表。"""
    if key in IOT_CLASSIC_ONLY_KEYS or key in CONTAINER_KEYS:
        return []
    # chartFrame=static 恒无数据绑定；未声明 chartFrame 时（可选字段）由 dataset 是否存在决定，
    # 与前端 ChartData/index.vue 的 isNotData 逻辑一致（frame === STATIC || option.dataset === undefined）
    if chart_frame == 'static':
        return []
    # composite 复合组件豁免 dataset 检查；其余组件需 option.dataset 存在才可绑定数据
    if chart_frame != 'composite' and not dataset_present:
        return []

    normalized = chart_key[1:] if chart_key.startswith('V') else chart_key
    if not normalized:
        return ['property', 'deviceStatus', 'deviceInfo']

    # property 恒为支持（源码：queryType === 'property' 直接返回 true）
    result = ['property']
    if normalized in LINE_CHART_KEYS:
        pass  # 折线组仅 property（且仅 log 模式）
    elif normalized in PIE_CHART_KEYS:
        result.append('deviceStatus')
    elif normalized in DEVICE_INFO_CHART_KEYS:
        result.extend(['deviceStatus', 'deviceInfo'])
    elif (normalized in TEXT_CHART_KEYS or normalized in OPEN_LIST_CHART_KEYS
          or normalized in RANK_LIST_CHART_KEYS or normalized in BAR_CHART_KEYS):
        if normalized in RANK_LIST_CHART_KEYS:
            result.append('deviceInfo')
        else:
            result.append('deviceStatus')
    elif normalized in RADAR_CHART_KEYS:
        result.append('deviceStatus')
    elif normalized in PROGRESS_CHART_KEYS:
        result.append('deviceStatus')
    else:
        result.extend(['deviceStatus', 'deviceInfo'])
    return result


def parse_regular_component(entry, enum_maps_by_package, global_enum_maps):
    """解析常规组件 index.ts，输出 components.json 条目（package/category/chartFrame 输出真实枚举值）。"""
    text = entry['text']
    package = entry['package']
    enum_maps = enum_maps_by_package.get(package, {})
    category_member = extract_enum_member(text, 'category', ('ChatCategoryEnum', 'PresetCategoryEnum'))
    category_name_member = extract_enum_member(text, 'categoryName', ('ChatCategoryEnumName', 'PresetCategoryNameEnum'))
    # category 输出枚举值（画布存档中实际存储的字符串，如 'Bars'），缺省回退成员名
    category = enum_maps.get('ChatCategoryEnum', {}).get(
        category_member,
        enum_maps.get('PresetCategoryEnum', {}).get(category_member, category_member))
    category_name = enum_maps.get('ChatCategoryEnumName', {}).get(
        category_name_member,
        enum_maps.get('PresetCategoryNameEnum', {}).get(category_name_member, category_name_member))
    package_member = extract_enum_member(text, 'package', ('PackagesCategoryEnum',))
    package_value = global_enum_maps.get('PackagesCategoryEnum', {}).get(package_member, package_member or package)
    chart_frame_member = extract_enum_member(text, 'chartFrame', ('ChartFrameEnum',))
    chart_frame = global_enum_maps.get('ChartFrameEnum', {}).get(chart_frame_member, chart_frame_member)

    config_ts = os.path.join(entry['dir'], 'config.ts')
    dataset_present = has_dataset(config_ts)

    key = extract_literal(text, 'key')
    chart_key = extract_literal(text, 'chartKey')
    return {
        'key': key,
        'chartKey': chart_key,
        'conKey': extract_literal(text, 'conKey'),
        'title': extract_literal(text, 'title'),
        'category': category,
        'categoryName': category_name,
        'package': package_value,
        'chartFrame': chart_frame,
        'image': extract_literal(text, 'image'),
        'iotQueryTypes': compute_iot_query_types(key, chart_key, chart_frame, dataset_present),
        'singleValueSelfRender': chart_key in IOT_SINGLE_VALUE_SELF_RENDER_CHARTS,
    }


def parse_presets(src, enum_maps_by_package, global_enum_maps):
    """解析 Presets 各子分类 index.ts：基础信息来自 Decorates/BorderCharts 底层 config，包/分类/标题/图片由 Preset 覆盖。"""
    results = []
    presets_dir = os.path.join(src, 'Presets')
    enum_maps = enum_maps_by_package.get('Presets', {})
    if not os.path.isdir(presets_dir):
        return results
    for sub in sorted(os.listdir(presets_dir)):
        sub_index = os.path.join(presets_dir, sub, 'index.ts')
        if not os.path.isfile(sub_index):
            continue
        text = read_text(sub_index)
        # import XxxConfig from '.../config'
        imports = {}
        for m in re.finditer(r"import\s+(\w+)\s+from\s+'([^']+)'", text):
            spec = m.group(2)
            if spec.endswith('/config'):
                imports[m.group(1)] = os.path.normpath(os.path.join(os.path.dirname(sub_index), spec + '.ts'))
        # export const PresetXxx: ConfigType = { ... }
        for m in re.finditer(r'export\s+const\s+(\w+):\s*ConfigType\s*=\s*\{(.*?)\n\}', text, re.S):
            block = m.group(2)
            spread = re.search(r'new\s+(\w+)\(\)', block)
            base_file = imports.get(spread.group(1)) if spread else None
            base_text = read_text(base_file) if base_file and os.path.isfile(base_file) else ''
            key = extract_literal(base_text, 'key')
            chart_key = extract_literal(base_text, 'chartKey')
            category_member = extract_enum_member(block, 'category', ('PresetCategoryEnum',))
            # category 输出枚举值（如 'Metrics'），缺省回退成员名
            category = enum_maps.get('PresetCategoryEnum', {}).get(category_member, category_member)
            category_name_member = extract_enum_member(block, 'categoryName', ('PresetCategoryNameEnum',))
            category_name = enum_maps.get('PresetCategoryNameEnum', {}).get(category_name_member, category_name_member)
            package_member = extract_enum_member(block, 'package', ('PackagesCategoryEnum',))
            package_value = global_enum_maps.get('PackagesCategoryEnum', {}).get(package_member, package_member or 'Presets')
            chart_frame_member = extract_enum_member(block, 'chartFrame', ('ChartFrameEnum',)) or \
                extract_enum_member(base_text, 'chartFrame', ('ChartFrameEnum',))
            chart_frame = global_enum_maps.get('ChartFrameEnum', {}).get(chart_frame_member, chart_frame_member)
            dataset_present = has_dataset(base_file) if base_file else False
            results.append({
                'key': key,
                'chartKey': chart_key,
                'conKey': extract_literal(base_text, 'conKey'),
                'title': extract_literal(block, 'title'),
                'category': category,
                'categoryName': category_name,
                'package': package_value,
                'chartFrame': chart_frame,
                'image': extract_literal(block, 'image'),
                'redirectComponent': extract_literal(block, 'redirectComponent'),
                'iotQueryTypes': compute_iot_query_types(key, chart_key, chart_frame, dataset_present),
                'singleValueSelfRender': chart_key in IOT_SINGLE_VALUE_SELF_RENDER_CHARTS,
            })
    # 保持 Presets/index.ts 的导出顺序：Metrics, Status, Compare, Common, Container
    root_order = []
    root_index = os.path.join(presets_dir, 'index.ts')
    if os.path.isfile(root_index):
        root_text = read_text(root_index)
        root_order = re.findall(r"import\s+\w+\s+from\s+'\./(\w+)'", root_text)
    order_map = {name: i for i, name in enumerate(root_order)}
    results.sort(key=lambda item: order_map.get(item['category'], 99))
    return results


def parse_iot_config_fields(type_ts_path):
    """从 chartEditStore.type.ts 提取 RequestIoTDeviceConfigType 字段表（字段名/类型/含义）。"""
    text = read_text(type_ts_path)
    start = text.index('export interface RequestIoTDeviceConfigType')
    brace_start = text.index('{', start)
    depth = 0
    end = brace_start
    for i in range(brace_start, len(text)):
        if text[i] == '{':
            depth += 1
        elif text[i] == '}':
            depth -= 1
            if depth == 0:
                end = i
                break
    body = text[brace_start + 1:end]

    fields = []
    pending_comments = []
    nested_prefix = ''
    for raw_line in body.splitlines():
        line = raw_line.strip()
        if not line:
            continue
        if line.startswith('//'):
            pending_comments.append(line.lstrip('/').strip())
            continue
        if re.match(r'^\}[,;]?$', line):
            if nested_prefix:
                nested_prefix = ''
                continue
            continue
        nested = re.match(r"^(\w+)\??:\s*\{\s*$", line)
        if nested:
            nested_prefix = nested.group(1) + '.'
            desc = '；'.join(pending_comments)
            fields.append((nested.group(1), 'object', desc or '嵌套配置对象'))
            pending_comments = []
            continue
        m = re.match(r"^(\w+)(\?)?\s*:\s*(.+?);?\s*(?://\s*(.*))?$", line)
        if m:
            name = nested_prefix + m.group(1)
            ftype = re.sub(r'\s+', ' ', m.group(3)).rstrip(';').strip()
            inline = m.group(4) or ''
            desc_parts = pending_comments + ([inline] if inline else [])
            desc = '；'.join(p for p in desc_parts if p)
            # 去掉源码注释中的 ── 分隔装饰
            desc = desc.replace('──', '').strip(' ；')
            fields.append((name, ftype, desc))
        pending_comments = []
    return fields


def note_for(item):
    """生成组件的“状态与注意点”列。"""
    notes = []
    if item['key'] in CONTAINER_KEYS:
        notes.append('分组容器（chartKey=group），不参与数据绑定')
    if item['key'] in IOT_CLASSIC_ONLY_KEYS:
        notes.append('经典数据面板限定 静态/AJAX，走专用数据配置，不走通用 IoT 面板')
    frame = item['chartFrame']
    if frame == 'static' and not notes:
        notes.append('无数据绑定')
    if not frame and not notes:
        notes.append('未声明 chartFrame，按 dataset 判定数据绑定')
    if item['singleValueSelfRender']:
        notes.append('单值自渲染，IoT 数据不走 ECharts dataset 覆盖')
    if item.get('redirectComponent'):
        notes.append('复合预置组件，redirectComponent → %s' % item['redirectComponent'])
    return '；'.join(notes) if notes else '—'


def scenario_for(item):
    return SCENARIOS.get(item['key']) or PACKAGE_SCENARIO_DEFAULT.get(item.get('package', ''), '—')


def iot_display(item):
    return ', '.join(item['iotQueryTypes']) if item['iotQueryTypes'] else '—'


def escape_md(value):
    return str(value).replace('|', '\\|')


def gen_json(components, snapshot_date):
    payload = {
        'snapshotDate': snapshot_date,
        'source': 'apps/web/packages/bigscreen/src/goview/packages/components',
        'components': [
            {
                'key': c['key'],
                'chartKey': c['chartKey'],
                'conKey': c['conKey'],
                'title': c['title'],
                'category': c['category'],
                'categoryName': c['categoryName'],
                'package': c['package'],
                'chartFrame': c['chartFrame'],
                'image': c['image'],
                'iotQueryTypes': c['iotQueryTypes'],
                'singleValueSelfRender': c['singleValueSelfRender'],
            }
            for c in components
        ],
    }
    return json.dumps(payload, ensure_ascii=False, indent=2) + '\n'


def gen_markdown(components, preset_items, snapshot_date, iot_fields, src):
    lines = []
    lines.append('# GoView 大屏组件清单')
    lines.append('')
    lines.append('> 快照日期：%s' % snapshot_date)
    lines.append('> 源码路径：`%s`（仓库内相对路径 `apps/web/packages/bigscreen/src/goview/packages/components`）' % src)
    lines.append('> 本文件由 `scripts/gen-view-components.py` 自动生成，请勿手工编辑；前端组件变更后重新运行脚本即可刷新。')
    lines.append('> 表中 package/category/chartFrame 为画布存档 JSON 实际存储的枚举值字符串（如 `Charts`/`Bars`/`echarts`）。')
    lines.append('')

    by_package = {}
    for c in components:
        by_package.setdefault(c['package'], []).append(c)

    # ── 总览 ──
    lines.append('## 总览')
    lines.append('')
    lines.append('| 分类 | 组件数 | 说明 |')
    lines.append('| --- | --- | --- |')
    total = 0
    for pkg in PACKAGE_ORDER:
        items = by_package.get(pkg, [])
        total += len(items)
        remark = {
            'Icons': '另有 127 个图标库动态条目（uim:/line-md:/wi: 前缀），统一重定向到 Icon 组件',
            'Photos': '无静态组件，资源库/本地上传/共享图片运行时动态生成，重定向到 Informations/Mores/Image',
            'Presets': '业务复合组件，通过 redirectComponent 指向底层实现',
        }.get(pkg, '')
        lines.append('| %s %s | %d | %s |' % (pkg, PACKAGE_CN[pkg], len(items), remark or '—'))
    lines.append('| **合计** | **%d** | 静态注册组件总数 |' % total)
    lines.append('')

    # ── 各分类明细 ──
    for pkg in PACKAGE_ORDER:
        items = by_package.get(pkg, [])
        lines.append('## %s %s（%d）' % (pkg, PACKAGE_CN[pkg], len(items)))
        lines.append('')
        if pkg == 'Icons':
            lines.append('Icons 分类只有一个真实组件 `Icon`（chartFrame=static）。图标面板中的条目由 '
                         '`Icons/MaterialLine`、`Icons/Common`、`Icons/Weather` 三个 index.ts 基于 IconConfig '
                         '动态生成（共 26 + 68 + 33 = 127 个），条目的 `image`/`dataset` 为图标名（`line-md:`/`uim:`/`wi:` 前缀），'
                         '`redirectComponent` 统一指向 `Icons/Default/Icon`。')
            lines.append('')
        if pkg == 'Photos':
            lines.append('Photos 分类没有静态组件：资源库（Library）条目运行时从资源库 API 加载，'
                         '本地（Local）条目来自用户上传并持久化在 localStorage，共享（Share）为远程图片列表；'
                         '三类条目均基于 `Informations/Mores/Image` 的 ImageConfig 生成，`chartFrame=static`，'
                         '`redirectComponent` 指向 `Informations/Mores/Image`。')
            lines.append('')
        if not items:
            continue
        lines.append('| key | chartKey | 标题 | chartFrame | 支持的 IoT queryType | 适用场景 | 状态与注意点 |')
        lines.append('| --- | --- | --- | --- | --- | --- | --- |')
        for c in items:
            lines.append('| %s | %s | %s | %s | %s | %s | %s |' % (
                c['key'], c['chartKey'], escape_md(c['title']), c['chartFrame'] or '（未声明）',
                iot_display(c), escape_md(scenario_for(c)), escape_md(note_for(c))))
        lines.append('')

    # ── 附录 A：5 文件结构 ──
    lines.append('## 附录 A：组件目录结构')
    lines.append('')
    lines.append('每个常规组件目录固定包含 5 个文件：')
    lines.append('')
    lines.append('| 文件 | 职责 |')
    lines.append('| --- | --- |')
    lines.append('| `index.ts` | ConfigType 元数据（key/chartKey/conKey/title/category/categoryName/package/chartFrame/image） |')
    lines.append('| `config.ts` | 默认配置类（attr/option/events 等，option 内含默认 dataset） |')
    lines.append('| `index.vue` | 画布渲染组件（按 chartKey 注册） |')
    lines.append('| `config.vue` | 右侧配置面板组件（按 conKey 注册） |')
    lines.append('| `data.json` | 静态示例数据（静态数据源模式使用） |')
    lines.append('')
    lines.append('例外：Icons/Photos 条目动态生成，无独立目录；Presets 通过 `redirectComponent` 复用底层实现组件的 5 文件。')
    lines.append('')

    # ── 附录 B：新增组件步骤 ──
    lines.append('## 附录 B：前端新增组件步骤（简要）')
    lines.append('')
    lines.append('1. 在对应包/子分类下新建组件目录，补齐 `index.ts` / `config.ts` / `index.vue` / `config.vue` / `data.json` 五个文件。')
    lines.append('2. 在子分类 `index.ts`（如 `Charts/Bars/index.ts`）中 import 并加入默认导出数组，完成注册。')
    lines.append('3. 新增子分类时同步维护包目录 `index.type.ts` 的 `ChatCategoryEnum` / `ChatCategoryEnumName` 枚举。')
    lines.append('4. 需要支持 IoT 数据源时，按组件形态更新 `ChartDataIoTDevice/queryTypeSupport.ts` 中的组件集合'
         '（折线/饼/柱/文字/表格/雷达/进度等分组），单值自渲染组件同步更新 `useChartDataFetch.hook.ts` 的 '
         '`IOT_SINGLE_VALUE_SELF_RENDER_CHARTS`。')
    lines.append('5. 重新运行 `scripts/gen-view-components.py` 刷新 components.json 与本清单。')
    lines.append('')

    # ── 附录 C：RequestIoTDeviceConfigType 字段表 ──
    lines.append('## 附录 C：RequestIoTDeviceConfigType 字段表')
    lines.append('')
    lines.append('来源：`apps/web/packages/bigscreen/src/goview/store/modules/chartEditStore/chartEditStore.type.ts`。')
    lines.append('')
    lines.append('| 字段名 | 类型 | 含义 |')
    lines.append('| --- | --- | --- |')
    for name, ftype, desc in iot_fields:
        lines.append('| `%s` | `%s` | %s |' % (name, escape_md(ftype), escape_md(desc or '—')))
    lines.append('')

    # ── 附录 D：IoT queryType 支持矩阵 ──
    lines.append('## 附录 D：IoT queryType 支持矩阵（按组件分组）')
    lines.append('')
    lines.append('| 组件分组 | 成员 chartKey | property | deviceStatus | deviceInfo | 备注 |')
    lines.append('| --- | --- | --- | --- | --- | --- |')
    matrix_rows = [
        ('折线组', 'LineCommon / LineGradients / LineGradientSingle / LineLinearSingle', True, False, False, '仅属性历史时序（log 模式）'),
        ('饼图组', 'PieCircle / PieCommon', True, True, False, '属性仅 latest 快照，适合组成占比'),
        ('柱状组', 'BarCommon / BarCrossrange / BarLine', True, True, False, '数值分类对比'),
        ('文字组', 'TextCommon / FlipperNumber', True, True, False, '单值文本 / 翻牌器'),
        ('表格-信息组', 'TablesBasic / TableList', True, True, True, '支持设备信息字段与状态统计'),
        ('表格-滚动组', 'TableScrollBoard', True, True, False, '滚动表格'),
        ('雷达组', 'Radar', True, True, False, '多维数值对比'),
        ('进度组', 'Process / WaterPolo', True, True, False, '单值比例，属性仅 latest'),
        ('胶囊组', 'CapsuleChart', True, True, True, 'queryType 默认全放行；属性仅 latest 模式（多设备同属性对比）'),
        ('其他可绑定数据组件', '其余 chartFrame 非 static 且有 dataset 的组件', True, True, True, '源码 isQueryTypeSupported 默认全部放行'),
        ('静态/装饰组件', 'chartFrame=static 或无 dataset', False, False, False, 'iotQueryTypes 为空数组'),
    ]
    for name, keys, prop, status, info, remark in matrix_rows:
        mark = lambda v: '✓' if v else '✗'
        lines.append('| %s | %s | %s | %s | %s | %s |' % (name, keys, mark(prop), mark(status), mark(info), remark))
    lines.append('')
    lines.append('补充说明：')
    lines.append('- `property` 查询类型在源码中对所有可绑定数据的组件恒为支持；折线组仅支持 log（历史时序）模式，饼图/胶囊/进度组仅支持 latest（最新值）模式。')
    lines.append('- `singleValueSelfRender=true` 的组件（VPieCircle / VProcess / VWaterPolo / VDial）拿到 IoT 单值后自行渲染，不走 ECharts dataset 覆盖。')
    lines.append('- BorderDeviceStatus / BorderRunningTimer 在经典数据面板中被限定为 静态/AJAX，不开放通用 IoT 面板。')
    lines.append('- 注意：源码 `getQueryTypeHint` 中 TextCommon 的提示语称“支持设备信息字段”，但 `isQueryTypeSupported` 实际判定文字组禁用 deviceInfo，两处不一致；本清单以后者（运行时实际判定）为准。')
    lines.append('')
    return '\n'.join(lines)


def main():
    parser = argparse.ArgumentParser(description='盘点 GoView 组件，生成 components.json 与 components.md')
    parser.add_argument('--src', default='/home/ubuntu/saas/apps/web/packages/bigscreen/src/goview/packages/components',
                        help='goview 组件根目录')
    parser.add_argument('--json-out', default='/home/ubuntu/saas/.gits/cli/cmd/view/data/components.json',
                        help='components.json 输出路径')
    parser.add_argument('--md-out', default='/home/ubuntu/saas/.gits/cli/skill/ur-view/references/components.md',
                        help='components.md 输出路径')
    parser.add_argument('--snapshot-date', default='2026-08-03', help='快照日期（YYYY-MM-DD）')
    args = parser.parse_args()

    src = os.path.abspath(args.src)
    if not os.path.isdir(src):
        print('组件目录不存在: %s' % src, file=sys.stderr)
        sys.exit(1)

    # 各包枚举表
    enum_maps_by_package = {}
    for pkg in PACKAGE_ORDER:
        enum_maps_by_package[pkg] = parse_enum_maps(os.path.join(src, pkg, 'index.type.ts'))

    # 全局枚举表（packages/index.type.ts）：PackagesCategoryEnum / ChartFrameEnum 成员名 → 枚举值，
    # components.json 的 package/category/chartFrame 输出枚举值（与大屏画布存档中实际存储的字符串一致）
    global_enum_maps = parse_enum_maps(os.path.join(os.path.dirname(src), 'index.type.ts'))

    # 常规组件（含 Icons 的 Icon；Presets 走专用解析）
    components = []
    for entry in find_registered_components(src):
        if entry['package'] == 'Presets':
            continue
        components.append(parse_regular_component(entry, enum_maps_by_package, global_enum_maps))

    # Presets 复合组件
    preset_items = parse_presets(src, enum_maps_by_package, global_enum_maps)

    # 按包顺序合并（Presets 在 Photos 之后、Interact 之前）
    ordered = []
    for pkg in PACKAGE_ORDER:
        if pkg == 'Presets':
            ordered.extend(preset_items)
        else:
            ordered.extend([c for c in components if c['package'] == pkg])
    components = ordered

    # IoT 配置字段表
    goview_root = os.path.dirname(os.path.dirname(src))
    type_ts = os.path.join(goview_root, 'store/modules/chartEditStore/chartEditStore.type.ts')
    iot_fields = parse_iot_config_fields(type_ts)

    json_text = gen_json(components, args.snapshot_date)
    md_text = gen_markdown(components, preset_items, args.snapshot_date, iot_fields, src)

    for out_path, content in ((args.json_out, json_text), (args.md_out, md_text)):
        os.makedirs(os.path.dirname(out_path), exist_ok=True)
        with open(out_path, 'w', encoding='utf-8') as f:
            f.write(content)
        print('已生成: %s' % out_path)

    # 控制台统计
    counts = {}
    for c in components:
        counts[c['package']] = counts.get(c['package'], 0) + 1
    print('组件总数: %d' % len(components))
    for pkg in PACKAGE_ORDER:
        print('  %s: %d' % (pkg, counts.get(pkg, 0)))


if __name__ == '__main__':
    main()
