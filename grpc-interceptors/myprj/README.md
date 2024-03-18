brew isntall protobuf

sudo go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
sudo go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

protoc --go_out=. --go-grpc_out=. ./pkg/proto/book.proto


For client
----------
brew install grpcui
grpcui -plaintext localhost:9901


brew install grpcurl

### grpc gateway
https://github.com/grpc-ecosystem/grpc-gateway

go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
    

protoc -I ./pkg/proto --go_out=./pkg/proto --go-grpc_out=./pkg/proto --grpc-gateway_out=. --go_opt paths=source_relative --go-grpc_opt paths=source_relative ./pkg/proto/book.proto

