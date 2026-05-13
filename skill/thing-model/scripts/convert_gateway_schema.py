#!/usr/bin/env python3
"""
网关导入模板 Excel → 通用物模型 CommonSchemaInfo JSON 转换脚本

用途:
    将 DLT645电表.xlsx、CJT188水表.xlsx 等网关导入模板转换为
    联犀平台通用物模型批量导入格式（CommonSchemaInfo JSON）

用法:
    python3 convert_gateway_schema.py

输出:
    schema.json - 通用物模型导入文件
    schema.zip  - 批量导入 ZIP 包

关键格式说明:
    1. CommonSchemaInfo.name 和 affordance 使用 StringValue 包装:
       - name: {"value": "中文名称"}
       - affordance: {"value": json.dumps(affordance_obj)}

    2. 批量导入 API 需要 ZIP 包，内含 schema.json

    3. 导入流程：
       a) 上传 ZIP 到 OSS: /api/v1/system/common/upload-file
       b) 调用批量导入: /api/v1/things/schema/common/batch-import
"""

import json
import pandas as pd
from typing import List, Dict, Any

# 单位映射（根据标识符前缀或名称关键词推断）
UNIT_MAP = {
    "U": "V",      # 电压
    "I": "A",      # 电流
    "P": "kW",     # 有功功率
    "Q": "kVar",   # 无功功率
    "S": "kVA",    # 视在功率
    "PF": "",      # 功率因数（无单位）
    "Freq": "Hz",  # 频率
    "EP": "kWh",   # 有功电能
    "EQ": "kWh",   # 无功电能
    "Flux": "m³",  # 流量
}

def get_unit(identifier: str, name: str) -> str:
    """根据标识符和名称推断单位"""
    name_lower = name.lower()

    # 根据名称关键词匹配（优先）
    if "电压" in name:
        return "V"
    elif "电流" in name:
        return "A"
    elif "有功功率" in name:
        return "kW"
    elif "无功功率" in name:
        return "kVar"
    elif "视在功率" in name:
        return "kVA"
    elif "功率因数" in name or "功率因素" in name:
        return ""
    elif "频率" in name:
        return "Hz"
    elif "电能" in name:
        return "kWh"
    elif "流量" in name:
        return "m³"

    # 根据标识符前缀匹配
    for prefix, unit in UNIT_MAP.items():
        if identifier.startswith(prefix):
            return unit

    return ""

def get_mode(access_mode: int) -> str:
    """转换读写模式: 0=R, 1=W, 2=RW"""
    if access_mode == 0:
        return "r"
    return "rw"

def get_step(data_type: str, decimals: int) -> str:
    """根据小数位生成步长"""
    if data_type == "double":
        if decimals <= 0:
            return "1"
        return "0." + "0" * (decimals - 1) + "1" if decimals > 0 else "1"
    return "1"

def convert_dlt645(filepath: str) -> List[Dict[str, Any]]:
    """转换 DLT645 电表 Excel

    Excel 列顺序（注意名称/标识符反直觉）:
        - "名称" 列 → 英文标识符（如 Ua, Ia）
        - "标识符" 列 → 中文名称（如 A相电压）
        - "数据类型" 列 → double/uint32/int32
        - "小数位" 列 → 整数

    输出 CommonSchemaInfo 格式:
        - type: 1 (property)
        - identifier: 英文标识符
        - name: {"value": 中文名称}  # StringValue 包装
        - affordance: {"value": json.dumps(affordance_obj)}  # StringValue 包装
    """
    df = pd.read_excel(filepath)

    # 跳过第一行（标题行）
    df = df.iloc[1:]

    schemas = []
    for _, row in df.iterrows():
        # 注意：Excel 中 "名称" 列是英文标识符，"标识符" 列是中文名称
        identifier = str(row["名称"]).strip()  # 英文标识符
        name = str(row["标识符"]).strip()  # 中文名称
        data_type = str(row["数据类型"]).strip().lower()
        decimals = int(row["小数位"]) if pd.notna(row["小数位"]) else 0

        # 数据类型映射
        if data_type == "double":
            define_type = "float"
            step = get_step(data_type, decimals)
        elif data_type == "uint32":
            define_type = "int"
            step = "1"
        elif data_type == "int32":
            define_type = "int"
            step = "1"
        else:
            define_type = "float"
            step = "0.001"

        unit = get_unit(identifier, name)

        # 特殊处理 Relay（跳合闸控制）- 使用 enum 类型
        if identifier == "Relay":
            affordance = {
                "isUseShadow": False,
                "isNoRecord": False,
                "define": {
                    "type": "enum",
                    "mapping": {"0": "合闸", "1": "跳闸"}
                },
                "mode": "rw"
            }
            user_perm = 3  # 可读写
        else:
            # 普通属性 - 使用 float/int 类型
            affordance = {
                "isUseShadow": False,
                "isNoRecord": False,
                "define": {
                    "type": define_type,
                    "min": "0",
                    "max": "999999999",
                    "step": step,
                },
                "mode": "r"  # 电表属性默认只读
            }
            if unit:
                affordance["define"]["unit"] = unit
            user_perm = 1  # 只读

        schema = {
            "type": 1,  # property
            "identifier": identifier,
            "name": {"value": name},  # StringValue 包装
            "desc": {},
            "required": 2,
            "isCanSceneLinkage": 1,
            "funcGroup": 1,
            "userPerm": user_perm,
            "affordance": {"value": json.dumps(affordance, ensure_ascii=False)},  # StringValue 包装
            "extendConfig": "{}",
            "order": 1,
            "controlMode": 1
        }

        schemas.append(schema)

    return schemas

def convert_cjt188(filepath: str) -> List[Dict[str, Any]]:
    """转换 CJT188 水表 Excel

    Excel 列顺序:
        - "功能名称" 列 → 英文标识符
        - "标识符" 列 → 中文名称
        - "读写类型(0:R,1:W,2:RW)" 列 → 0/1/2
        - "数据类型(0:uint32,1:int32,2:double)" 列 → 0/1/2
        - "小数位" 列 → 整数
    """
    df = pd.read_excel(filepath)

    # 跳过第一行（标题行）
    df = df.iloc[1:]

    schemas = []
    for _, row in df.iterrows():
        identifier = str(row["功能名称"]).strip()  # 英文标识符
        name = str(row["标识符"]).strip()  # 中文名称
        access_mode = int(row["读写类型(0:R,1:W,2:RW)"]) if pd.notna(row["读写类型(0:R,1:W,2:RW)"]) else 0
        data_type = int(row["数据类型(0:uint32,1:int32,2:double)"]) if pd.notna(row["数据类型(0:uint32,1:int32,2:double)"]) else 2
        decimals = int(row["小数位"]) if pd.notna(row["小数位"]) else 0

        mode = get_mode(access_mode)

        # 数据类型映射: 0:uint32, 1:int32, 2:double
        if data_type == 2:
            define_type = "float"
            step = get_step("double", decimals)
        else:
            define_type = "int"
            step = "1"

        # 特殊处理阀门状态和控制
        if identifier == "ValveStatus":
            affordance = {
                "isUseShadow": False,
                "isNoRecord": False,
                "define": {
                    "type": "enum",
                    "mapping": {"0": "开", "1": "关", "2": "异常"}
                },
                "mode": "r"
            }
            user_perm = 1
        elif identifier == "ValveCtrl":
            affordance = {
                "isUseShadow": False,
                "isNoRecord": False,
                "define": {
                    "type": "enum",
                    "mapping": {"0": "开阀", "1": "关阀"}
                },
                "mode": "rw"
            }
            user_perm = 3
        else:
            unit = get_unit(identifier, name)
            affordance = {
                "isUseShadow": False,
                "isNoRecord": False,
                "define": {
                    "type": define_type,
                    "min": "0",
                    "max": "99999999",
                    "step": step,
                },
                "mode": mode
            }
            if unit:
                affordance["define"]["unit"] = unit
            user_perm = 3 if access_mode in [1, 2] else 1

        schema = {
            "type": 1,  # property
            "identifier": identifier,
            "name": {"value": name},  # StringValue 包装
            "desc": {},
            "required": 2,
            "isCanSceneLinkage": 1,
            "funcGroup": 1,
            "userPerm": user_perm,
            "affordance": {"value": json.dumps(affordance, ensure_ascii=False)},  # StringValue 包装
            "extendConfig": "{}",
            "order": 1,
            "controlMode": 1
        }

        schemas.append(schema)

    return schemas

def main():
    """主函数：读取 Excel、转换、生成 ZIP 包"""
    # Excel 文件路径（相对于 docs 目录）
    dlt645_path = "docs/应用/能耗管理/网关导入模板/DLT645电表.xlsx"
    cjt188_path = "docs/应用/能耗管理/网关导入模板/CJT188水表.xlsx"

    all_schemas = []

    # 转换 DLT645 电表
    print("转换 DLT645 电表...")
    dlt_schemas = convert_dlt645(dlt645_path)
    print(f"  - 生成 {len(dlt_schemas)} 个属性")
    all_schemas.extend(dlt_schemas)

    # 转换 CJT188 水表
    print("转换 CJT188 水表...")
    cjt_schemas = convert_cjt188(cjt188_path)
    print(f"  - 生成 {len(cjt_schemas)} 个属性")
    all_schemas.extend(cjt_schemas)

    # 输出到文件
    output_path = "/tmp/schema_import/schema.json"
    with open(output_path, "w", encoding="utf-8") as f:
        json.dump(all_schemas, f, ensure_ascii=False, indent=2)

    print(f"\n总计 {len(all_schemas)} 个物模型属性")
    print(f"输出文件: {output_path}")

    # 创建 ZIP 包（批量导入需要）
    import zipfile
    zip_path = "/tmp/schema_import/schema.zip"
    with zipfile.ZipFile(zip_path, "w", zipfile.ZIP_DEFLATED) as zf:
        zf.write(output_path, "schema.json")
    print(f"ZIP 包: {zip_path}")

if __name__ == "__main__":
    main()