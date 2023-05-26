


gen:
	@go run github.com/99designs/gqlgen generate
	@go generate ./ent

protoc:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

account-server:
	@go run grpc/account-service/main.go
