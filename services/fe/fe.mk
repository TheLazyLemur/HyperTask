#### FE ####

fe_generate:
	@cd services/fe && \
		templ generate

fe_run: fe_generate
	@go run ./services/fe/...

fe_build: fe_generate
	@go build -o bin/fe ./services/fe/internal/server/...

