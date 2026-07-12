all: help

help: ## Print this help message
	@grep -E '^[a-zA-Z._-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run-go
run-go: ## Run the Go service
	cd golang && go run ./cmd/server

.PHONY: run-python
run-python: ## Run the Python service
	cd python && uv sync && uv run main.py

.PHONY: run-ts
run-ts: ## Run the TypeScript service
	cd ts && npm install && npm run build && npm start

.PHONY: test
test: test-go test-python test-ts ## Run all test suites

.PHONY: test-go
test-go: ## Run Go tests
	cd golang && go test ./... -race -cover

.PHONY: test-python
test-python: ## Run Python tests
	cd python && uv sync && uv run pytest

.PHONY: test-ts
test-ts: ## Run TypeScript tests
	cd ts && npm install && npm test
