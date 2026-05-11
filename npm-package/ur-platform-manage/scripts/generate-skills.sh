#!/usr/bin/env bash
set -euo pipefail
exec /usr/local/bin/ur-platform-manage generate-skills "$@"
