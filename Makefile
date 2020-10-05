all: agent cli proxy gateway rm um jm dm-master dm-node

protos:
	protoc \
		-I. \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		--go_out=plugins=grpc,paths=source_relative:. \
		--grpc-gateway_out=logtostderr=true,paths=source_relative:. \
		--swagger_out=logtostderr=true,allow_merge=true:docs \
		api/proto/*.proto

docs:
	mv docs/apidocs.swagger.json docs/swagger.json

pkger: docs
	pkger

proxy: protos
	go build -o build ./internal/reverse_proxy

gateway: pkger
	go build -o build ./cmd/gateway

agent: protos
	go build -o build ./cmd/agent

rm: protos
	go build -o build ./internal/resource_manager

um: protos
	go build -o build ./internal/user_manager

jm: protos
	go build -o build ./cmd/job_manager

dm-master: pkger
	go build -o build ./internal/data_manager/master

dm-node: pkger
	go build -o build ./internal/data_manager/node

cli: protos
	go build -o build ./internal/cli

clean:
	rm -f api/proto/*.go
