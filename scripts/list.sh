#!/usr/bin/env bash

HTTP="http --verify=no --check-status"

$HTTP "$INSTANCE_HOST/jobs" "Authorization: Bearer $ACCESS_TOKEN"
