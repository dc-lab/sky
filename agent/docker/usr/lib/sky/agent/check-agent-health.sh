#!/bin/sh

REFERENCE=$(cat /var/run/sky/agent/health.info)
sleep 0.5
COMPARABLE=$(cat /var/run/sky/agent/health.info)

if [ "$REFERENCE" == "$COMPARABLE" ]; then
  echo "Agent hasn't updated health.info for a long time"
  exit 1
else
  echo "Agent seems healthy"
fi
