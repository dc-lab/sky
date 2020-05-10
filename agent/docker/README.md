### Rebuilding the docker container

The main Agent container used for cloud integration lives here. Rebuild it with
```
cd ~/sky/agent
docker build -t sky/agent-inside-cloud -f ./docker/Dockerfile .
```
