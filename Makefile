data_manager_protos:
	protoc -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc,paths=source_relative:. api/proto/data_manager/*.proto
	protoc -I. -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true,paths=source_relative:. api/proto/data_manager/*.proto
