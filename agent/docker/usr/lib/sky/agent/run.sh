#!/bin/sh

mkdir -p /var/log/supervisor/children
mkdir -p /var/log/sky/agent

echo "$SKY_RM_TOKEN" > /etc/sky/agent/token
if [ "$SKY_AGENT_LOG_STDOUT" = "1" ]; then
  ln -s /dev/stdout /var/log/sky/agent/agent.log
fi

exec /usr/bin/supervisord -c /etc/supervisor/supervisord.conf
