#!/bin/bash
# 通用物模型批量导入脚本
#
# 用法:
#   ./batch_import_schema.sh <server_url> [token]
#
# 示例:
#   ./batch_import_schema.sh http://115.190.3.202:7777
#   ./batch_import_schema.sh http://1.13.180.134:7777 <token>
#
# 前提:
#   1. 已生成 schema.json 和 schema.zip（运行 convert_gateway_schema.py）
#   2. 有目标服务器的登录 token

set -e

SERVER_URL="${1:-http://localhost:7777}"
TOKEN="${2:-}"
ZIP_PATH="/tmp/schema_import/schema.zip"

# 如果没有提供 token，自动登录获取
if [ -z "$TOKEN" ]; then
    echo "正在登录获取 token..."
    LOGIN_RESP=$(curl -s -X POST "${SERVER_URL}/api/v1/system/user/self/login" \
        -H "app-id: 1" \
        -H "Content-Type: application/json" \
        -d '{"loginType":"pwd","account":"administrator","password":"iThings666","pwdType":1}')

    # 提取 accessToken
    TOKEN=$(echo "$LOGIN_RESP" | python3 -c "
import json, sys
data = json.load(sys.stdin)
if data.get('code') == 0:
    print(data.get('data', {}).get('accessToken', ''))
else:
    print(f'ERROR: {data}')
    sys.exit(1)
")

    if [ -z "$TOKEN" ] || [[ "$TOKEN" == ERROR* ]]; then
        echo "登录失败: $LOGIN_RESP"
        exit 1
    fi
    echo "Token: ${TOKEN:0:50}..."
fi

# 步骤1: 上传 ZIP 到 OSS
echo "正在上传 ZIP 文件到 OSS..."
UPLOAD_RESP=$(curl -s -X POST "${SERVER_URL}/api/v1/system/common/upload-file" \
    -H "token: $TOKEN" \
    -F "file=@${ZIP_PATH}")

# 提取 fileUri（注意：使用 fileUri 而非 filePath）
FILE_URI=$(echo "$UPLOAD_RESP" | python3 -c "
import json, sys
data = json.load(sys.stdin)
if data.get('code') == 0:
    # fileUri 包含完整路径如 /oss/ithings-temporary/xxx.zip
    print(data.get('data', {}).get('fileUri', ''))
else:
    print(f'ERROR: {data}')
    sys.exit(1)
")

if [ -z "$FILE_URI" ] || [[ "$FILE_URI" == ERROR* ]]; then
    echo "上传失败: $UPLOAD_RESP"
    exit 1
fi
echo "OSS 文件: $FILE_URI"

# 步骤2: 调用批量导入 API
echo "正在批量导入物模型..."
IMPORT_RESP=$(curl -s -X POST "${SERVER_URL}/api/v1/things/schema/common/batch-import" \
    -H "token: $TOKEN" \
    -H "Content-Type: application/json" \
    -d "{\"schemasFile\": \"$FILE_URI\"}")

# 解析导入结果
echo "$IMPORT_RESP" | python3 -c "
import json, sys
data = json.load(sys.stdin)
if data.get('code') == 0:
    result = data.get('data', {})
    print(f'导入成功:')
    print(f'  - 成功: {result.get(\"succCount\", 0)}')
    print(f'  - 错误: {result.get(\"errCount\", 0)}')
    print(f'  - 忽略: {result.get(\"ignoreCount\", 0)}')
    if result.get('errSchemas'):
        print(f'  - 错误详情:')
        for err in result.get('errSchemas', []):
            print(f'    {err}')
else:
    print(f'导入失败: {data}')
"

echo "完成!"