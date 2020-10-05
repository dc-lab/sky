export RP_HTTP_ADDRESS=:18080
export RP_LOGS_DIR=~/hse/third/sky/sky/logs
export RP_DB_USER=postgres
export RP_DB_PASSWORD=db_pasword
export RP_DB_HOST=localhost:5432
export RP_DB_NAME=sky_users
export RP_DB_SSL=false

export UM_HTTP_ADDRESS=:18081
export UM_GRPC_ADDRESS=:18082
export UM_LOGS_DIR=~/hse/third/sky/sky/logs
export UM_DB_USER=postgres
export UM_DB_PASSWORD=db_password
export UM_DB_HOST=localhost:5432
export UM_DB_NAME=sky_users
export UM_DB_SSL=false

export RM_HTTP_ADDRESS=:18083
export RM_GRPC_ADDRESS=:18084
export RM_DM_ADDRESS=localhost:18086
export RM_LOGS_DIR=~/hse/third/sky/sky/logs
export RM_DB_USER=postgres
export RM_DB_PASSWORD=db_password
export RM_DB_HOST=localhost:5432
export RM_DB_NAME=sky_rm
export RM_DB_SSL=false

export SKY_NO_LOGFILES=1
export SKY_RM_ADDR=resource_manager:18084
export SKY_RM_TOKEN=ABCD

export DM_MASTER_HTTP_BIND_ADDRESS=:18085
export DM_MASTER_GRPC_BIND_ADDRESS=:18086
export DM_MASTER_POSTGRES_ADDRESS="host=localhost port=5432 user=postgres password=db_password dbname=sky_files sslmode=disable"

export DM_NODE_HTTP_BIND_ADDRESS=:18087
export DM_NODE_ACCESS_ADDRESS=localhost:18080
export DM_NODE_MASTER_ADDRESS=localhost:18086
export DM_NODE_PUSH_INTERVAL=1s
export DM_NODE_STORAGE_DIR=/tmp/node1/storage
