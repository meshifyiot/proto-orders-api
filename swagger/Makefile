.PHONY: requirements all orders-api

GRPC_GW_PATH := $(shell go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/)/..
GOOGLEAPIS_PATH := "$(GRPC_GW_PATH)/third_party/googleapis"
PROTOS := $(shell cd ../proto && find orders/*.proto)

all:
	@echo
	@echo "[Generate swagger documentation...]"
	protoc -I=$(GOOGLEAPIS_PATH) -I=${GRPC_GW_PATH} -I=../proto \
	--swagger_out=logtostderr=true,use_go_templates=true,allow_merge=true,merge_file_name=orders-api:orders ${PROTOS}