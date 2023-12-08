tidy:
	@go mod tidy

vet:
	@go vet ./...

test: vet
	@go test ./...

clean:
	@rm -rf bin

include ./services/fe/fe.mk
include ./services/task/task.mk
