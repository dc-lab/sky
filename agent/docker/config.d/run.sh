#!/bin/sh

echo "$SKY_RM_TOKEN" > /usr/lib/sky/agent/lib/token

exec /usr/lib/sky/agent/lib/agent
