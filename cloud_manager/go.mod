module github.com/dc-lab/sky/cloud_manager

go 1.13

require (
	github.com/aws/aws-sdk-go v1.30.24
	github.com/dc-lab/sky/api/proto/cloud_manager v0.0.0-20200524150756-2184e08dcb44
	github.com/dc-lab/sky/api/proto/common v0.0.0-20200524150756-2184e08dcb44
	github.com/jackc/pgx/v4 v4.6.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.4.0
	google.golang.org/grpc v1.29.1
)
