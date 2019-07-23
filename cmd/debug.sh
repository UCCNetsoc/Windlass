#!/bin/bash

reset=$(ps -ef | grep 'windlass/__debug_bin' | grep -v grep | awk '{print $2}' | xargs -r kill)
eval $reset

reset=$(ps -ef | grep 'dlv' | grep -v grep | awk '{print $2}' | xargs -r kill)
eval $reset

echo "Starting Delve"
dlv debug ./cmd/windlass --build-flags '-mod vendor' -l 0.0.0.0:3456 --headless=true --api-version=2 &