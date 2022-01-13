#!/usr/bin/env bash

set -euo pipefail

ID=$1

$HTTP "$INSTANCE_HOST/jobs/$ID" "Authorization: Bearer $ACCESS_TOKEN"
