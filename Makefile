all: help

help: ## Print this help message
	@grep -E '^[a-zA-Z._-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: hello
hello: ## Print "Hello, World!"
	echo "Hello, World!"

.PHONY: run-go
run-go: ## Run the Go service
	cd golang && go run .

.PHONY: run-python
run-python: ## Run the Python service
	cd python && uv sync && uv run python main.py

.PHONY: run-ts
run-ts: ## Run the TypeScript service
	cd ts && npm install && npm run dev
