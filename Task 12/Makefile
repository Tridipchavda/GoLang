.PHONY: all proto clean

proto:
	protoc --go_out=. --go-grpc_out=. ./protobuf/*.proto && cd ./grpcService && go get