# -*- coding: utf-8 -*-
"""解析 DXF v3：全量实体分块解析（不分段，全局处理）"""
import sys, io, collections

sys.stdout = io.TextIOWrapper(sys.stdout.buffer, encoding='utf-8')

dxf_path = sys.argv[1] if len(sys.argv) > 1 else '1.dxf'
mode = sys.argv[2] if len(sys.argv) > 2 else 'all'

with open(dxf_path, 'r', encoding='utf-8', errors='replace') as f:
    lines = f.read().splitlines()

n = len(lines)
entities = []
cur = None
for i in range(0, n - 1, 2):
    code = lines[i].strip()
    val = lines[i+1]
    if code == '0':
        if cur: entities.append(cur)
        cur = {'type': val.strip(), 'c': collections.defaultdict(list)}
    elif cur is not None:
        try: cur['c'][int(code)].append(val)
        except ValueError: pass
if cur: entities.append(cur)

def get(e, k, idx=0, default=None):
    v = e['c'].get(k)
    return v[idx] if v and len(v) > idx else default

def fnum(s, default=None):
    try: return float(s)
    except (TypeError, ValueError): return default

def clean(t):
    # 去掉 MTEXT 控制码
    import re
    t = re.sub(r'\\[A-Za-z][^;\\]*;?', '', t)
    t = t.replace('\\P', ' / ').replace('{', '').replace('}', '')
    return t.strip()

def txt_of(e):
    return clean(''.join(e['c'].get(3, [])) + (e['c'][1][0] if e['c'].get(1) else ''))

if mode in ('all', 'stats'):
    print("=== 图层分布（全部实体） ===")
    layers = collections.Counter(get(e, 8, 0, '?') for e in entities)
    for l, c in layers.most_common():
        print(f"  {l}: {c}")

    print("\n=== TEXT/MTEXT/ATTRIB 全部文字（按图层） ===")
    for e in entities:
        if e['type'] in ('TEXT', 'MTEXT', 'ATTRIB'):
            t = txt_of(e)
            if t:
                layer = get(e, 8, 0, '?')
                x = fnum(get(e, 10)); y = fnum(get(e, 20))
                h = fnum(get(e, 40), 0)
                print(f"[{layer}] ({x:.0f},{y:.0f}) h={h:.0f} : {t}")

    print("\n=== INSERT 块引用（含坐标） ===")
    for e in entities:
        if e['type'] == 'INSERT':
            name = get(e, 2, 0, '?')
            x = fnum(get(e, 10)); y = fnum(get(e, 20))
            layer = get(e, 8, 0, '?')
            print(f"[{layer}] {name} @({x:.0f},{y:.0f})")

    print("\n=== LWPOLYLINE / LINE 按图层统计 ===")
    geo = collections.Counter()
    for e in entities:
        if e['type'] in ('LWPOLYLINE', 'LINE', 'POLYLINE'):
            geo[(e['type'], get(e, 8, 0, '?'))] += 1
    for (t, l), c in geo.most_common():
        print(f"  {t} [{l}]: {c}")
