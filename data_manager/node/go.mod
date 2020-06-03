module github.com/dc-lab/sky/data_manager/node

go 1.14

require (
	github.com/dc-lab/sky/api/proto/data_manager v0.0.0-20200511235702-bae58568a911
	github.com/dc-lab/sky/data_manager/common v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi v4.1.1+incompatible
	github.com/go-chi/render v1.0.1
	github.com/joho/godotenv v1.3.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.6.3
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/net v0.0.0-20200501053045-e0ff5e5a1de5 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200212174721-66ed5ce911ce // indirect
	google.golang.org/grpc v1.27.1
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/dc-lab/sky/api/proto/data_manager => /home/sergey/hse/third/sky/sky/api/proto/data_manager

replace github.com/dc-lab/sky/data_manager/common => /home/sergey/hse/third/sky/sky/data_manager/common