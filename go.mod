module github.com/dc-lab/sky

go 1.15

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/cenkalti/backoff/v4 v4.0.2
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v1.4.2-0.20200213202729-31a86c4ab209
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/golang-migrate/migrate/v4 v4.12.2
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.15.0
	github.com/jackc/pgx/v4 v4.9.0
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.8.0
	github.com/markbates/pkger v0.15.1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pressly/lg v1.1.1
	github.com/shirou/gopsutil v2.20.8+incompatible
	github.com/sirupsen/logrus v1.6.0
	github.com/snowzach/rotatefilehook v0.0.0-20180327172521-2f64f265f58c
	github.com/spf13/cobra v1.0.0
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.5.1
	github.com/struCoder/pidusage v0.1.3
	github.com/swaggo/http-swagger v0.0.0-20200308142732-58ac5e232fba
	github.com/swaggo/swag v1.6.7
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	google.golang.org/genproto v0.0.0-20200925023002-c2d885f95484
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gorm.io/driver/postgres v1.0.2
	gorm.io/gorm v1.20.2
)

replace ./api/proto => ./generated/api/proto
