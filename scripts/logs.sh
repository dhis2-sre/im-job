#!/usr/bin/env bash

HTTP="http --verify=no --check-status"

GROUP_ID=$1
RUN_ID=$2

echo "{
  \"groupId\": $GROUP_ID
  }
}" | $HTTP --stream "$INSTANCE_HOST/jobs/running/$RUN_ID/logs" "Authorization: Bearer $ACCESS_TOKEN"
