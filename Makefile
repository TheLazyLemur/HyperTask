clean:
	@rm -rf bin

#### TASK ####

task_generate:
	@cd services/task && \
		protoc --proto_path=internal/proto internal/proto/*.proto --go_out=. --go-grpc_out=.

task_run: task_generate
	@go run ./services/task/...

task_build: task_generate
	@go build -o bin/task ./services/task/internal/server/...

#### FE ####

fe_generate:
	@cd services/fe && \
		templ generate

fe_run: fe_generate
	@go run ./services/fe/...

fe_build: fe_generate
	@go build -o bin/fe ./services/fe/internal/server/...

