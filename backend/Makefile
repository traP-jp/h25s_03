.PHONY: tools
tools:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: dev
dev:
	go mod download
	go generate ./...
	export $(shell grep -v '^#' .env.dev | xargs) && go run .