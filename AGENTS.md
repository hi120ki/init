# Repository Guidelines

## Project Structure & Module Organization

This repository hosts language-specific starters that all mirror the same pipeline: load environment configuration, validate it, then log through a structured JSON logger. `golang/` is the reference implementation, a single-module service split into `cmd/server` (entry point), `internal/env` (configuration), and `internal/logger` (logging). `python/` is a uv-managed app mirroring that layout with `main.py`, `env.py`, and `logger.py`. `ts/` mirrors it with `src/index.ts`, `src/env.ts`, and `src/logger.ts`, compiled into `dist/` by `tsc`. Keep assets inside their language directory and place tests next to the code they exercise (`golang/...`, `python/tests/`, `ts/src/__tests__/`).

## Build, Test, and Development Commands

- `cd golang && go run ./cmd/server` starts the Go binary with structured logging.
- `cd python && uv sync && uv run main.py` installs dependencies, then runs the Python variant.
- `cd ts && npm install && npm run dev` runs the TypeScript service in watch mode via `tsx`; `npm run build && npm start` runs the compiled output; `npm run typecheck` type-checks without emitting; `npm run clean` removes `dist/`.
- `make help` at the repository root lists shared Make targets (`run-*` and `test-*` per language, `test` for all); update it when adding cross-project workflows.

## Coding Style & Naming Conventions

The Go implementation is the master: when behavior or structure diverges across languages, align Python and TypeScript to Go, not the other way around. Follow idiomatic patterns per language: keep Go package names lowercase and run `go fmt ./...` before committing. Python stays type-annotated, PEP 8 compliant, and snake_case; configuration goes through `pydantic-settings` and logging through `structlog`. TypeScript stays in ES module syntax with camelCase; configuration is validated with `zod` and logging goes through `pino`. Avoid TS enums (`erasableSyntaxOnly` is enabled); use `as const` objects instead. Keep shared environment keys such as `ENVIRONMENT` and `PORT` identical everywhere.

Formatting and linting are enforced by prek hooks (`prek.toml`): `golangci-lint` for Go, `ruff` (check + format) for Python, `oxfmt` for TypeScript/JSON/Markdown/YAML, plus `typos` and `gitleaks` repo-wide.

## Testing Guidelines

Add unit tests whenever you modify behavior. Go tests belong in the same package and run with `cd golang && go test ./...`. Python tests live under `python/tests/` and run on `pytest` with `cd python && uv sync && uv run pytest`. TypeScript tests live under `ts/src/__tests__/` and run on the Node.js test runner (via `tsx`) with `cd ts && npm test`. Missing runnable tests for new features will block review.

## Commit & Pull Request Guidelines

History currently uses short imperative subjects (“Initial commit”); keep that tone or adopt Conventional Commits if it clarifies scope. Each PR should note the language surface touched, setup steps, and config changes, and include logs or console output when behavior shifts. Link issues when available and call out new environment variables or Make targets so reviewers can verify locally. CI runs a per-language build-and-test workflow when files under that language's directory change, plus CodeQL, dependency review, and Scorecard checks.

## Environment & Configuration Tips

All starters read an optional `.env` file (see the `.env.example` files) or exported variables. `ENVIRONMENT` (`development` or `production`, default `development`) controls log verbosity—debug in development, info in production. `PORT` (default `8080`) sets the listen port. Keep secrets in deployment tooling rather than the repo, never log them, and document any new flags or configuration fields you introduce.
