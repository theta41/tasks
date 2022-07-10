.PHONY: proto, swagger

all: proto swagger test build

test:
	 go tool cover -func cover.out

build:
	go mod download
	go build -o main cmd/main.go

swagger:
	swag init -g internal/http/api.go

proto:
	protoc -I proto/ --go_out=api --go-grpc_out=api proto/auth.proto