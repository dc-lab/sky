### Rebuilding the docker container

The main Agent container used for cloud integration lives here. Rebuild it with
```
cd ~/sky/agent
docker build -t sky/agent -f ./docker/Dockerfile .
```

### Running the docker container

```
docker run --network=host -e SKY_RM_TOKEN=docker-token sky/agent
```