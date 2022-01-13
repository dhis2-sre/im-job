#!/usr/bin/env bash

HTTP="http --verify=no --check-status"

GROUP_NAME=$1
RUN_ID=$2

GROUP_ID=$($HTTP --check-status "$INSTANCE_HOST/groups-name-to-id/$GROUP_NAME" "Authorization: Bearer $ACCESS_TOKEN")

echo "{
  \"groupId\": $GROUP_ID
  }
}" | $HTTP "$INSTANCE_HOST/jobs/running/$RUN_ID/status" "Authorization: Bearer $ACCESS_TOKEN"
