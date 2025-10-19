import logging
import os
import sys

import structlog
from dotenv import load_dotenv
from pydantic import Field
from pydantic_settings import BaseSettings


class Config(BaseSettings):
    api_key: str = Field(..., validation_alias="API_KEY")


def main():
    environment = os.getenv("ENVIRONMENT")
    log_level = logging.DEBUG if environment != "production" else logging.INFO

    structlog.configure(
        processors=[
            structlog.stdlib.filter_by_level,
            structlog.stdlib.add_log_level,
            structlog.stdlib.add_logger_name,
            structlog.processors.TimeStamper(fmt="iso"),
            structlog.processors.JSONRenderer(),
        ],
        logger_factory=structlog.stdlib.LoggerFactory(),
        wrapper_class=structlog.stdlib.BoundLogger,
        cache_logger_on_first_use=True,
    )
    logging.basicConfig(
        format="%(message)s",
        stream=sys.stdout,
        level=log_level,
    )
    log = structlog.get_logger()

    load_dotenv()

    try:
        config = Config()
    except Exception as e:
        log.error("Failed to load configuration", error=str(e))
        sys.exit(1)

    log.info("Application started", api_key=config.api_key)


if __name__ == "__main__":
    main()
