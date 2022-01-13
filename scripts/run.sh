#!/usr/bin/env bash

HTTP="http --verify=no --check-status"

ID=$1
JOB_TYPE=$2
GROUP_NAME=$3
INSTANCE_NAME=$4

GROUP_ID=$($HTTP --check-status "$INSTANCE_HOST/groups-name-to-id/$GROUP_NAME" "Authorization: Bearer $ACCESS_TOKEN")
INSTANCE_ID=$($HTTP --check-status "$INSTANCE_HOST/instances-name-to-id/$GROUP_ID/$INSTANCE_NAME" "Authorization: Bearer $ACCESS_TOKEN")

echo "{
  \"jobType\": \"$JOB_TYPE\",
  \"groupId\": $GROUP_ID,
  \"targetId\": $INSTANCE_ID,
  \"payload\": {
    \"KEY\": \"VAL\"
  }
}" | $HTTP "$INSTANCE_HOST/jobs/$ID/run" "Authorization: Bearer $ACCESS_TOKEN"
