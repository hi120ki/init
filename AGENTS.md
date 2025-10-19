# Repository Guidelines

## Project Structure & Module Organization

This repository hosts language-specific starters. `golang/` provides a single-module service wired with `slog`, `godotenv`, and `envconfig`. `python/` is a uv-managed app whose `main.py` mirrors the Go pipeline with `pydantic-settings`. `ts/` contains a TypeScript entry at `src/index.ts`, compiled into `dist/` by `tsc`. Keep assets inside their language directory and place tests next to the code they exercise (`golang/...`, `python/tests/`, `ts/src/__tests__/`).

## Build, Test, and Development Commands

- `cd golang && go run .` starts the Go binary with structured logging.
- `cd python && uv sync && uv run python main.py` installs dependencies, then runs the Python variant.
- `cd ts && npm install && npm run dev` builds the TypeScript bundle and runs the generated Node service.
- `make help` lists shared Make targets; update it when adding cross-project workflows.

## Coding Style & Naming Conventions

Follow idiomatic patterns: run `go fmt ./...` before committing Go changes and keep package names lowercase. Python code stays type-annotated, PEP 8 compliant, and uses snake_case; log with `structlog`. TypeScript remains in ES module syntax, uses camelCase, and validates configuration with `zod`. Keep shared environment keys such as `API_KEY` and `ENVIRONMENT` identical everywhere.

## Testing Guidelines

Add unit tests whenever you modify behavior. Go tests belong in the same package and run with `cd golang && go test ./...`. Python tests live under `python/tests/`; add `pytest` to `pyproject.toml` and execute with `cd python && uv run pytest`. TypeScript tests go under `ts/src/__tests__/`; wire them into `npm test` so the placeholder command stops failing. Missing runnable tests for new features will block review.

## Commit & Pull Request Guidelines

History currently uses short imperative subjects (“Initial commit”); keep that tone or adopt Conventional Commits if it clarifies scope. Each PR should note the language surface touched, setup steps, and config changes, and include logs or console output when behavior shifts. Link issues when available and call out new environment variables or Make targets so reviewers can verify locally.

## Environment & Configuration Tips

All starters expect `.env` files or exported variables providing an `API_KEY`; keep secrets in deployment tooling rather than the repo. Log verbosity is controlled by `ENVIRONMENT`—default to non-production verbosity while developing and document any new flags or configuration fields you introduce.
