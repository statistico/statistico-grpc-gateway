build:
	protoc -I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
       --go_out proto/gen --go_opt paths=source_relative \
       --go-grpc_out proto/gen --go-grpc_opt paths=source_relative \
       --grpc-gateway_out proto/gen --grpc-gateway_opt paths=source_relative \
       --openapiv2_out proto/gen --openapiv2_opt logtostderr=true --openapiv2_opt allow_merge=true,merge_file_name=statistico \
       --proto_path=proto \
       proto/*.proto
