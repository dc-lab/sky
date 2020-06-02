module github.com/dc-lab/sky/cloud_manager

go 1.13

require (
	github.com/aws/aws-sdk-go v1.31.8
	github.com/cenkalti/backoff/v4 v4.0.2
	github.com/dc-lab/sky/api/proto/cloud_manager v0.0.0-20200602003057-28be1fed3712
	github.com/dc-lab/sky/api/proto/common v0.0.0-20200602003057-28be1fed3712
	github.com/dc-lab/sky/api/proto/resource_manager v0.0.0-20200602003057-28be1fed3712
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/google/uuid v1.1.1
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.0
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
	google.golang.org/grpc v1.29.1
)
