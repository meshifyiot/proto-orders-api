.PHONY: all orders-api

GRPC_GW_PATH := $(shell go list -f '{{ .Dir }}' github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/)/..
GOOGLEAPIS_PATH := "$(GRPC_GW_PATH)/third_party/googleapis"
PROTOS := $(shell cd ../proto && find orders/*.proto)

all:
	@echo
	@echo "[Generate go source files...]"
	protoc -I=$(GOOGLEAPIS_PATH) -I=${GRPC_GW_PATH} -I=../proto --go_out=plugins=grpc:. ${PROTOS}