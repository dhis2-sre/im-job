#!/usr/bin/env bash

HTTP="http --verify=no --check-status"

GROUP_ID=$1
RUN_ID=$2

echo "{
  \"groupId\": $GROUP_ID
  }
}" | $HTTP "$INSTANCE_HOST/jobs/running/$RUN_ID/status" "Authorization: Bearer $ACCESS_TOKEN"