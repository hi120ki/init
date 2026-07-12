# Python Starter

Mirrors the Go starter pipeline: load configuration from environment variables (with an optional `.env` file), validate it, then log through a structured JSON logger whose level follows the environment.

## Usage

```sh
uv sync          # install dependencies
uv run main.py   # run the application
uv run pytest    # run tests
```

## Configuration

| Variable      | Default       | Description                   |
| ------------- | ------------- | ----------------------------- |
| `ENVIRONMENT` | `development` | `development` or `production` |
| `PORT`        | `8080`        | Port number (1-65535)         |
