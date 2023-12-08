generate:
	@protoc --proto_path=internal/proto internal/proto/*.proto --go_out=. --go-grpc_out=.

run:
	@go run ./...

build:
	@go build -o bin/task ./internal/server/...

clean:
	@rm -rf bin
