#!/usr/bin/env bash

set -euo pipefail

GROUP_NAME=$1
RUN_ID=$2

GROUP_ID=$($HTTP --check-status "$INSTANCE_HOST/groups-name-to-id/$GROUP_NAME" "Authorization: Bearer $ACCESS_TOKEN")

echo "{
  \"groupId\": $GROUP_ID
  }
}" | $HTTP --stream "$INSTANCE_HOST/jobs/running/$RUN_ID/logs" "Authorization: Bearer $ACCESS_TOKEN"
