### Rebuilding the docker container

The main Agent container used for cloud integration lives here. Rebuild it with

```sh
cd ~/sky/agent
docker build -t sky/agent -f ./docker/Dockerfile .
```

### Running the docker container

```sh
docker run -dit --rm --network=host --name sky-agent -e SKY_RM_TOKEN=docker-token sky/agent
```

### Attaching to the docker container

```sh
docker exec -ti sky-agent /bin/sh
```

### Interact with supervisor

```sh
/ # supervisord ctl status
agent:agent                      RUNNING   pid 16, uptime 0:01:04
```

```sh
/ # supervisord ctl stop agent
agent: stopped
/ # supervisord ctl status
agent:agent                      EXITED    2020-05-16 22:56:44.757649749 +0000 UTC m=+248.722554191
```

```sh
/ # supervisord ctl start agent
agent: started
/ # supervisord ctl status
agent:agent                      RUNNING   pid 147, uptime 0:00:02
```

### File structure

#### In git
```text
/etc -- for configs
/usr -- for runtime
```

#### In docker container
```text
/etc/sky/agent/config.json -- config for agent
/etc/sky/agent/token -- token for agent
/etc/supervisor/supervisord.conf -- config for supervisord

/usr/bin/sky-agent -- agent binary
/usr/bin/supervisord -- supervisord binary

/usr/lib/sky/agent/run.sh -- launcher for agent

/var/log/sky/agent/agent.log -- logs of sky-agent
/var/log/supervisor/supervisord.log.* -- logs of supervisord
/var/log/supervisor/children/agent.stdout.log.* -- stdout of agent
/var/log/supervisor/children/agent.stderr.log.* -- stderr of agent

/var/run/supervisord.pid -- PID of supervisord

/var/tmp/sky/agent -- workdir for agent
```
