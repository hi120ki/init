"""Application entry point, mirroring the Go starter pipeline."""

from structlog.typing import FilteringBoundLogger

import sys

from pydantic import ValidationError

import env
from logger import new_logger


def main() -> None:
    try:
        cfg: env.Env = env.load()
    except ValidationError as error:
        print(f"Failed to load environment variables: {error}", file=sys.stderr)
        sys.exit(1)

    logger: FilteringBoundLogger = new_logger(cfg.environment)

    logger.info(event="application started", port=cfg.port)


if __name__ == "__main__":
    main()
