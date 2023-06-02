


gen:
	@go run github.com/99designs/gqlgen generate
	@go generate ./ent

protoc:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	@export PATH="$PATH:$(go env GOPATH)/bin"
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

account-server:
	@go run grpc/account-service/main.go
