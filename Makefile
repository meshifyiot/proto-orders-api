.PHONY: all build go swagger swagger-ui typescript clean

PROTOC_GEN_GO_VERSION=1.4.2
PROTOC_GEN_GRPC_GATEWAY_VERSION=1.14.7
PROTOC_GEN_SWAGGER_VERSION=1.14.7

SWAGGER_DIR=`pwd`/swagger/orders
SWAGGER_PORT=3010

all: clean build go swagger typescript

build:
	docker-compose build
	
go:
	docker-compose run orders-api-go

swagger:
	docker-compose run orders-api-swagger

swagger-ui:
	@echo "If chrome did not open after 2s you can go to http://localhost:${SWAGGER_PORT} to see the docs"
	@(sleep 2s ; open 'http://localhost:${SWAGGER_PORT}') &
	@docker run -p ${SWAGGER_PORT}:8080 -e SWAGGER_JSON=/tmp/orders-api.swagger.json -v ${SWAGGER_DIR}:/tmp swaggerapi/swagger-ui

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