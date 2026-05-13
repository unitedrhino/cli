#!/usr/bin/env python3
"""
场景联动 JSON 结构校验工具

用法:
    python3 validate_scene.py <file>
    cat scene.json | python3 validate_scene.py -
    python3 validate_scene.py --template auto
    python3 validate_scene.py --template manual
"""

import json
import sys
from typing import Any

# 枚举值定义（与 Go 代码保持一致）
SCENE_TYPES = {"auto", "manual"}
DEVICE_MODES = {"single", "multi"}
STATUSES = {1, 2}
TRIGGER_TYPES = {"device", "timer", "weather"}
SELECT_TYPES = {"all", "fixed", "area", "areaWithChildren", "group"}
TRIGGER_DEVICE_MODES = {"edge", "level"}
TRIGGER_DEVICE_TYPES = {"connected", "disConnected", "propertyReport", "eventReport"}
EXEC_TYPES = {"at", "sunRises", "sunSet", "loop"}
REPEAT_TYPES = {"once", "week", "mount", "allDay", "customRange"}
WEATHER_TYPES = {"temp", "humidity"}
WHEN_RANGE_TYPES = {"date", "time"}
DATE_RANGE_TYPES = {"allDay", "workday", "weekend", "holiday", "customRange", "customWeek"}
TIME_RANGE_TYPES = {"allDay", "light", "night", "customRange"}
TERM_COND_TYPES = {"and", "or"}
TERM_COLUMN_TYPES = {"property", "event", "time", "weather"}
ACTION_TYPES = {"device", "delay", "notify", "alarm", "scene"}
ACTION_DEVICE_TYPES = {"propertyControl", "action"}
NOTIFY_TYPES = {"sms", "email", "dingTalk", "wxMini", "message", "phoneCall", "wxEWebhook"}
NOTIFY_CODES = {"ruleScene", "ruleDeviceAlarm"}
NOTIFY_USER_TYPES = {
    "account", "userID", "deviceOwner", "deviceProjectAdmin",
    "deviceAreaAdmin", "deviceProjectAll", "deviceAreaAll"
}
ALARM_MODES = {"trigger", "relieve"}
CMP_TYPES = {"eq", "not", "btw", "gt", "gte", "lt", "lte", "in", "all"}
STATE_KEEP_TYPES = {"duration", "repeat"}

CMP_VALUES_COUNT = {
    "eq": 1, "not": 1, "btw": 2, "gt": 1, "gte": 1,
    "lt": 1, "lte": 1, "in": -1, "all": 0
}


class Validator:
    def __init__(self):
        self.errors: list[str] = []
        self.warnings: list[str] = []

    def error(self, path: str, msg: str):
        self.errors.append(f"[{path}] {msg}")

    def warn(self, path: str, msg: str):
        self.warnings.append(f"[{path}] {msg}")

    def validate_enum(self, path: str, value: Any, allowed: set, required: bool = True):
        if value is None:
            if required:
                self.error(path, "必填字段缺失")
            return
        if value not in allowed:
            self.error(path, f"非法值 '{value}'，允许: {sorted(allowed)}")

    def validate_string(self, path: str, value: Any, required: bool = True):
        if value is None:
            if required:
                self.error(path, "必填字段缺失")
            return
        if not isinstance(value, str):
            self.error(path, f"应为 string 类型，实际为 {type(value).__name__}")

    def validate_int(self, path: str, value: Any, required: bool = True):
        if value is None:
            if required:
                self.error(path, "必填字段缺失")
            return
        if not isinstance(value, int):
            self.error(path, f"应为 int 类型，实际为 {type(value).__name__}")

    def validate_array(self, path: str, value: Any, required: bool = True):
        if value is None:
            if required:
                self.error(path, "必填字段缺失")
            return
        if not isinstance(value, list):
            self.error(path, f"应为 array 类型，实际为 {type(value).__name__}")

    def validate_info(self, data: dict, path: str = ""):
        if not isinstance(data, dict):
            self.error(path, "场景联动 JSON 应为 object 类型")
            return

        self.validate_enum(f"{path}.type", data.get("type"), SCENE_TYPES)
        self.validate_enum(f"{path}.deviceMode", data.get("deviceMode"), DEVICE_MODES)
        self.validate_string(f"{path}.name", data.get("name"))

        scene_type = data.get("type")
        device_mode = data.get("deviceMode")

        if device_mode == "single":
            self.validate_string(f"{path}.productID", data.get("productID"))
            self.validate_string(f"{path}.deviceName", data.get("deviceName"))

        status = data.get("status")
        if status is not None and status not in STATUSES:
            self.error(f"{path}.status", f"非法值 {status}，允许: {STATUSES}")

        if scene_type == "auto":
            if "if" not in data or data["if"] is None:
                self.error(f"{path}.if", "自动触发类型必填")
            else:
                self.validate_if(data.get("if"), f"{path}.if")

        if "when" in data and data["when"] is not None:
            self.validate_when(data.get("when"), f"{path}.when")

        if "then" not in data or data["then"] is None:
            self.error(f"{path}.then", "执行动作必填")
        else:
            self.validate_then(data.get("then"), f"{path}.then")

    def validate_if(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        triggers = data.get("triggers")
        self.validate_array(f"{path}.triggers", triggers)
        if isinstance(triggers, list):
            for i, t in enumerate(triggers):
                self.validate_trigger(t, f"{path}.triggers[{i}]")

    def validate_trigger(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_enum(f"{path}.type", data.get("type"), TRIGGER_TYPES)
        trigger_type = data.get("type")

        if trigger_type == "device":
            if "device" not in data or data["device"] is None:
                self.error(f"{path}.device", "设备触发类型必填")
            else:
                self.validate_trigger_device(data.get("device"), f"{path}.device")
        elif trigger_type == "timer":
            if "timer" not in data or data["timer"] is None:
                self.error(f"{path}.timer", "定时触发类型必填")
            else:
                self.validate_trigger_timer(data.get("timer"), f"{path}.timer")
        elif trigger_type == "weather":
            if "weather" not in data or data["weather"] is None:
                self.error(f"{path}.weather", "天气触发类型必填")
            else:
                self.validate_trigger_weather(data.get("weather"), f"{path}.weather")

    def validate_trigger_device(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_string(f"{path}.productID", data.get("productID"))
        self.validate_enum(f"{path}.selectType", data.get("selectType"), SELECT_TYPES)
        self.validate_enum(f"{path}.mode", data.get("mode"), TRIGGER_DEVICE_MODES, required=False)
        self.validate_enum(f"{path}.type", data.get("type"), TRIGGER_DEVICE_TYPES)

        select_type = data.get("selectType")
        if select_type == "fixed":
            self.validate_string(f"{path}.deviceName", data.get("deviceName"))

        device_type = data.get("type")
        if device_type in {"propertyReport", "eventReport"}:
            self.validate_string(f"{path}.dataID", data.get("dataID"))

        if "stateKeep" in data and data["stateKeep"] is not None:
            self.validate_state_keep(data.get("stateKeep"), f"{path}.stateKeep")

        # 校验 Compare
        has_term_type = "termType" in data and data["termType"]
        has_terms = "terms" in data and data["terms"]
        if has_term_type:
            self.validate_compare(data, path)
        elif has_terms:
            self.validate_cmps(data.get("terms"), f"{path}.terms")

    def validate_trigger_timer(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_enum(f"{path}.execType", data.get("execType"), EXEC_TYPES)
        self.validate_enum(f"{path}.repeatType", data.get("repeatType"), REPEAT_TYPES, required=False)

        exec_type = data.get("execType")
        if exec_type == "at":
            exec_at = data.get("execAt")
            if exec_at is not None:
                if not isinstance(exec_at, int) or exec_at < 0 or exec_at > 86400:
                    self.error(f"{path}.execAt", f"应在 0-86400 之间，实际: {exec_at}")
        elif exec_type in {"sunRises", "sunSet"}:
            exec_add = data.get("execAdd", 0)
            if not isinstance(exec_add, int) or exec_add < -10800 or exec_add > 10800:
                self.error(f"{path}.execAdd", f"应在 -10800~10800 之间，实际: {exec_add}")

    def validate_trigger_weather(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_enum(f"{path}.type", data.get("type"), WEATHER_TYPES)
        self.validate_compare(data, path)

    def validate_state_keep(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_enum(f"{path}.type", data.get("type"), STATE_KEEP_TYPES)
        self.validate_int(f"{path}.value", data.get("value"))

    def validate_when(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return

        if "validRanges" in data and data["validRanges"] is not None:
            self.validate_when_ranges(data.get("validRanges"), f"{path}.validRanges")
        if "invalidRanges" in data and data["invalidRanges"] is not None:
            self.validate_when_ranges(data.get("invalidRanges"), f"{path}.invalidRanges")
        if "conditions" in data and data["conditions"] is not None:
            self.validate_conditions(data.get("conditions"), f"{path}.conditions")

    def validate_when_ranges(self, data: list, path: str):
        self.validate_array(path, data)
        if isinstance(data, list):
            for i, r in enumerate(data):
                self.validate_when_range(r, f"{path}[{i}]")

    def validate_when_range(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_enum(f"{path}.type", data.get("type"), WHEN_RANGE_TYPES)
        range_type = data.get("type")
        if range_type == "date":
            self.validate_date_range(data.get("dateRange"), f"{path}.dateRange")
        elif range_type == "time":
            self.validate_time_range(data.get("timeRange"), f"{path}.timeRange")

    def validate_date_range(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_enum(f"{path}.type", data.get("type"), DATE_RANGE_TYPES)
        date_type = data.get("type")
        if date_type == "customRange":
            self.validate_string(f"{path}.startDate", data.get("startDate"))
            self.validate_string(f"{path}.endDate", data.get("endDate"))

    def validate_time_range(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_enum(f"{path}.type", data.get("type"), TIME_RANGE_TYPES)
        time_type = data.get("type")
        if time_type == "customRange":
            start = data.get("startTime")
            end = data.get("endTime")
            if start is not None and (not isinstance(start, int) or start < 0 or start > 86400):
                self.error(f"{path}.startTime", f"应在 0-86400 之间，实际: {start}")
            if end is not None and (not isinstance(end, int) or end < 0 or end > 86400):
                self.error(f"{path}.endTime", f"应在 0-86400 之间，实际: {end}")

    def validate_conditions(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_enum(f"{path}.type", data.get("type"), TERM_COND_TYPES, required=False)
        terms = data.get("terms")
        if terms is not None:
            self.validate_array(f"{path}.terms", terms, required=False)
            if isinstance(terms, list):
                for i, t in enumerate(terms):
                    self.validate_term(t, f"{path}.terms[{i}]")

    def validate_term(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_enum(f"{path}.columnType", data.get("columnType"), TERM_COLUMN_TYPES)
        column_type = data.get("columnType")
        if column_type == "property":
            if "property" in data and data["property"] is not None:
                self.validate_term_property(data.get("property"), f"{path}.property")
        elif column_type == "weather":
            if "weather" in data and data["weather"] is not None:
                self.validate_trigger_weather(data.get("weather"), f"{path}.weather")
        elif column_type == "time":
            if "time" in data and data["time"] is not None:
                self.validate_term_time(data.get("time"), f"{path}.time")

    def validate_term_property(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_string(f"{path}.productID", data.get("productID"))
        self.validate_string(f"{path}.deviceName", data.get("deviceName"))
        has_term_type = "termType" in data and data["termType"]
        has_terms = "terms" in data and data["terms"]
        if has_term_type:
            self.validate_compare(data, path)
        elif has_terms:
            self.validate_cmps(data.get("terms"), f"{path}.terms")

    def validate_term_time(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        time_types = {"sys", "sunRises", "sunSet"}
        self.validate_enum(f"{path}.type", data.get("type"), time_types)
        self.validate_compare(data, path)

    def validate_then(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        actions = data.get("actions")
        self.validate_array(f"{path}.actions", actions)
        if isinstance(actions, list):
            for i, a in enumerate(actions):
                self.validate_action(a, f"{path}.actions[{i}]")

    def validate_action(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_enum(f"{path}.type", data.get("type"), ACTION_TYPES)
        action_type = data.get("type")

        if action_type == "delay":
            self.validate_int(f"{path}.delay", data.get("delay"))
        elif action_type == "device":
            if "device" not in data or data["device"] is None:
                self.error(f"{path}.device", "设备动作类型必填")
            else:
                self.validate_action_device(data.get("device"), f"{path}.device")
        elif action_type == "notify":
            if "notify" not in data or data["notify"] is None:
                self.error(f"{path}.notify", "通知动作类型必填")
            else:
                self.validate_action_notify(data.get("notify"), f"{path}.notify")
        elif action_type == "alarm":
            if "alarm" not in data or data["alarm"] is None:
                self.error(f"{path}.alarm", "告警动作类型必填")
            else:
                self.validate_action_alarm(data.get("alarm"), f"{path}.alarm")
        elif action_type == "scene":
            if "scene" not in data or data["scene"] is None:
                self.error(f"{path}.scene", "场景动作类型必填")
            else:
                self.validate_action_scene(data.get("scene"), f"{path}.scene")

    def validate_action_device(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_string(f"{path}.productID", data.get("productID"))
        self.validate_enum(f"{path}.selectType", data.get("selectType"), SELECT_TYPES)
        self.validate_enum(f"{path}.type", data.get("type"), ACTION_DEVICE_TYPES)
        select_type = data.get("selectType")
        if select_type == "fixed":
            self.validate_string(f"{path}.deviceName", data.get("deviceName"))

        has_data_id = "dataID" in data and data["dataID"]
        has_value = "value" in data and data["value"]
        has_values = "values" in data and data["values"]

        if not has_data_id and not has_values:
            self.error(path, "设备动作需填写 dataID 或 values")

        if has_values:
            values = data.get("values")
            self.validate_array(f"{path}.values", values)
            if isinstance(values, list):
                for i, v in enumerate(values):
                    self.validate_device_value(v, f"{path}.values[{i}]")

    def validate_device_value(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_string(f"{path}.dataID", data.get("dataID"))
        self.validate_string(f"{path}.value", data.get("value"))

    def validate_action_notify(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_enum(f"{path}.type", data.get("type"), NOTIFY_TYPES)
        self.validate_enum(f"{path}.notifyCode", data.get("notifyCode"), NOTIFY_CODES)
        self.validate_enum(f"{path}.userType", data.get("userType"), NOTIFY_USER_TYPES)

        user_type = data.get("userType")
        if user_type == "account":
            self.validate_array(f"{path}.accounts", data.get("accounts"))
        elif user_type == "userID":
            self.validate_array(f"{path}.userIDs", data.get("userIDs"))

    def validate_action_alarm(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_enum(f"{path}.mode", data.get("mode"), ALARM_MODES)

    def validate_action_scene(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_int(f"{path}.sceneID", data.get("sceneID"))

    def validate_compare(self, data: dict, path: str):
        term_type = data.get("termType")
        if term_type is not None:
            self.validate_enum(f"{path}.termType", term_type, CMP_TYPES, required=False)
            if term_type in CMP_VALUES_COUNT:
                expected = CMP_VALUES_COUNT[term_type]
                values = data.get("values")
                if expected == 0:
                    if values:
                        self.warn(f"{path}.values", f"termType='all' 时 values 应为空或省略")
                elif expected == -1:
                    if not values or len(values) == 0:
                        self.error(f"{path}.values", f"termType='{term_type}' 需要至少1个值")
                elif values is None or len(values) != expected:
                    self.error(f"{path}.values", f"termType='{term_type}' 需要 {expected} 个值，实际: {len(values) if values else 0}")

    def validate_cmps(self, data: list, path: str):
        self.validate_array(path, data)
        if isinstance(data, list):
            for i, c in enumerate(data):
                self.validate_cmp(c, f"{path}[{i}]")

    def validate_cmp(self, data: dict, path: str):
        if not isinstance(data, dict):
            self.error(path, "应为 object 类型")
            return
        self.validate_string(f"{path}.column", data.get("column"))
        self.validate_compare(data, path)


def generate_template(scene_type: str) -> dict:
    """生成场景联动 JSON 模板"""
    base = {
        "type": scene_type,
        "deviceMode": "multi",
        "name": "",
        "desc": "",
        "status": 1,
        "when": {
            "validRanges": [],
            "invalidRanges": [],
            "conditions": {
                "type": "and",
                "terms": []
            }
        },
        "then": {
            "actions": []
        }
    }
    if scene_type == "auto":
        base["if"] = {
            "triggers": [
                {
                    "type": "device",
                    "order": 1,
                    "device": {
                        "productID": "",
                        "selectType": "fixed",
                        "deviceName": "",
                        "type": "propertyReport",
                        "dataID": "",
                        "termType": "gt",
                        "values": [""]
                    }
                }
            ]
        }
    return base


def main():
    if len(sys.argv) < 2:
        print("用法: validate_scene.py <file>       # 校验 JSON 文件")
        print("       cat file.json | validate_scene.py -")
        print("       validate_scene.py --template auto   # 生成自动触发模板")
        print("       validate_scene.py --template manual # 生成手动触发模板")
        sys.exit(2)

    arg = sys.argv[1]

    if arg == "--template" or arg == "-t":
        if len(sys.argv) < 3:
            print("错误: --template 需要参数 auto 或 manual")
            sys.exit(2)
        template_type = sys.argv[2]
        if template_type not in {"auto", "manual"}:
            print(f"错误: 模板类型应为 auto 或 manual，实际: {template_type}")
            sys.exit(2)
        tmpl = generate_template(template_type)
        print(json.dumps(tmpl, ensure_ascii=False, indent=2))
        sys.exit(0)

    # 读取输入
    if arg == "-":
        content = sys.stdin.read()
    else:
        try:
            with open(arg, "r", encoding="utf-8") as f:
                content = f.read()
        except FileNotFoundError:
            print(f"错误: 文件不存在: {arg}")
            sys.exit(1)
        except Exception as e:
            print(f"错误: 读取文件失败: {e}")
            sys.exit(1)

    # 解析 JSON
    try:
        data = json.loads(content)
    except json.JSONDecodeError as e:
        print(f"JSON 解析错误: {e}")
        sys.exit(1)

    # 校验
    v = Validator()
    v.validate_info(data)

    # 输出结果
    if v.errors:
        print(f"❌ 校验失败，共 {len(v.errors)} 个错误:")
        for e in v.errors:
            print(f"  {e}")
    else:
        print("✅ 校验通过")

    if v.warnings:
        print(f"⚠️  共 {len(v.warnings)} 个警告:")
        for w in v.warnings:
            print(f"  {w}")

    sys.exit(1 if v.errors else 0)


if __name__ == "__main__":
    main()
