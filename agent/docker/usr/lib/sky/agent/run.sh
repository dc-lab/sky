#!/bin/sh

mkdir -p /var/log/supervisor/children
mkdir -p /var/log/sky/agent

stdout_logfile="/var/log/supervisor/children/agent.stdout.log"
stdout_logfile_maxbytes="50MB"
stderr_logfile="/var/log/supervisor/children/agent.stderr.log"
stderr_logfile_maxbytes="50MB"

if [ "$SKY_NO_LOGFILES" = "1" ]; then
  ln -s /dev/stdout /var/log/sky/agent/agent.log
  stdout_logfile="/dev/stdout"
  stdout_logfile_maxbytes="0"
  stderr_logfile="/dev/stderr"
  stderr_logfile_maxbytes="0"
fi

cat /etc/sky/agent/config-template.json | sed "s#RESOURCE_MANAGER_ADDRESS#${SKY_RM_ADDR}#g" > /etc/sky/agent/config.json
cat /etc/supervisor/supervisord-template.conf | sed "s#STDOUT_LOGFILE_MAXBYTES#${stdout_logfile_maxbytes}#g" \
                                              | sed "s#STDERR_LOGFILE_MAXBYTES#${stderr_logfile_maxbytes}#g" \
                                              | sed "s#STDOUT_LOGFILE#${stdout_logfile}#g" \
                                              | sed "s#STDERR_LOGFILE#${stderr_logfile}#g" \
                                              > /etc/supervisor/supervisord.conf

echo "$SKY_RM_TOKEN" > /etc/sky/agent/token

exec /usr/bin/supervisord -c /etc/supervisor/supervisord.conf
