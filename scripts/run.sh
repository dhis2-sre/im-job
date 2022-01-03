#!/usr/bin/env bash

HTTP="http --verify=no --check-status"

ID=$1
GROUP_ID=$2

echo "{
  \"groupId\": $GROUP_ID,
  \"payload\": {
    \"KEY\": \"VAL\"
  }
}" | $HTTP "$INSTANCE_HOST/jobs/$ID/run" "Authorization: Bearer $ACCESS_TOKEN"
