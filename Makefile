.PHONY: all build go swagger typescript clean

PROTOC_GEN_GO_VERSION=1.4.2
PROTOC_GEN_GRPC_GATEWAY_VERSION=1.14.7
PROTOC_GEN_SWAGGER_VERSION=1.14.7

all: clean build go swagger typescript

build:
	docker-compose build
	
go:
	docker-compose run orders-api-go

swagger:
	docker-compose run orders-api-swagger

typescript:
	@echo "\n[Generate Typescript (based on swagger documentation)...]"
	docker-compose run orders-api-typescript
	./typescript/cleanup.sh

clean:
	rm -rf go/orders/*.go
	go mod tidy
	go get \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@v${PROTOC_GEN_GRPC_GATEWAY_VERSION} \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@v${PROTOC_GEN_SWAGGER_VERSION} \
		github.com/golang/protobuf/protoc-gen-go@v${PROTOC_GEN_GO_VERSION}
	rm -rf swagger/orders/*.json
	rm -rf typescript/orders/*