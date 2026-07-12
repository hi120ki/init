# Init

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fhi120ki%2Finit.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fhi120ki%2Finit?ref=badge_shield)

This repository is a Go / Python / TypeScript starter pack. Each language implements the same pipeline: load configuration from environment variables (with an optional `.env` file), validate it, then log through a structured JSON logger whose level follows the environment.

## Quick Start

- Go: `cd golang && go run ./cmd/server`
- Python: `cd python && uv sync && uv run main.py`
- TypeScript: `cd ts && npm install && npm run dev`

Run `make help` at the repository root for run and test targets covering all three languages.

## Configuration

| Variable      | Default       | Description                   |
| ------------- | ------------- | ----------------------------- |
| `ENVIRONMENT` | `development` | `development` or `production` |
| `PORT`        | `8080`        | Port number                   |

## Guidelines

See `AGENTS.md` for the full contributor guide, including structure, testing, and PR expectations.

## License

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fhi120ki%2Finit.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fhi120ki%2Finit?ref=badge_large)
