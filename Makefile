GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/dist
GOENVVARS := GOBIN=$(GOBIN)
GOBINARY := flight
GOCMD := $(GOBASE)/cmd

LINT := $$(go env GOPATH)/bin/golangci-lint run --timeout=5m -E whitespace -E gosec -E gci -E misspell -E gomnd -E gofmt -E goimports -E golint --exclude-use-default=false --max-same-issues 0
BUILD := $(GOENVVARS) go build $(LDFLAGS) -o $(GOBIN)/$(GOBINARY) $(GOCMD)

.PHONY: build
build: ## Builds the binary locally into ./dist
	$(BUILD)

.PHONY: lint
lint: ## Runs the linter
	$(LINT)

.PHONY: run
run:build ## Runs the flight server
	./dist/flight runserver --port 9090

.PHONY: test
test: ## Runs only short tests without checking race conditions
	go test --cover -short -p 1 ./...

.PHONY: proto-gen
proto-gen: ## Generates code from proto file 
	cd proto/flight/v1 && protoc --proto_path=. --go_out=../../../pathctrl/pb --go-grpc_out=../../../pathctrl/pb  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative path.proto

## Help display.
## Pulls comments from beside commands and prints a nicely formatted
## display with the commands and their usage information.
.DEFAULT_GOAL := help

.PHONY: help
help: ## Prints this help
		@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'