module github.com/dc-lab/sky/data_manager/master

go 1.14

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/cenkalti/backoff/v4 v4.0.2
	github.com/dc-lab/sky/api/proto/data_manager v0.0.0
	github.com/go-chi/render v1.0.1
	github.com/go-openapi/spec v0.19.7 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.5
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.4.0
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.6.3
	github.com/swaggo/swag v1.6.5
	golang.org/x/net v0.0.0-20200501053045-e0ff5e5a1de5 // indirect
	golang.org/x/tools v0.0.0-20200504152539-33427f1b0364 // indirect
	google.golang.org/grpc v1.27.1
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/dc-lab/sky/api/proto/data_manager => /home/sergey/hse/third/sky/sky/api/proto/data_manager
