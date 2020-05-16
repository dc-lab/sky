#!/bin/sh

mkdir -p /var/log/supervisor/children
mkdir -p /var/lib/supervisor

echo "$SKY_RM_TOKEN" > /etc/sky/agent/token

exec /usr/bin/supervisord -c /etc/supervisor/supervisord.conf
