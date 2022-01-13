#!/usr/bin/env bash

set -euo pipefail

$HTTP "$INSTANCE_HOST/jobs" "Authorization: Bearer $ACCESS_TOKEN"
