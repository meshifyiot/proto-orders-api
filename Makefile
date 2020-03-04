generate: generate_protoc generate_typescript

generate_protoc:
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		--go_out=plugins=grpc:./grpc \
		--swagger_out=logtostderr=true,allow_merge=true,merge_file_name=api:swagger \
		proto/*.proto

generate_typescript:
	./scripts/generate_typescript.sh
