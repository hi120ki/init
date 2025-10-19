import logging
import sys
from enum import Enum

import structlog
from dotenv import load_dotenv
from pydantic import Field, ValidationError
from pydantic_settings import BaseSettings
from structlog.stdlib import BoundLogger


class Environment(str, Enum):
    DEVELOPMENT = "development"
    PRODUCTION = "production"

    @property
    def log_level(self) -> int:
        if self is Environment.PRODUCTION:
            return logging.INFO
        return logging.DEBUG


class Config(BaseSettings):
    api_key: str = Field(..., validation_alias="API_KEY")
    environment: Environment = Field(
        default=Environment.DEVELOPMENT, validation_alias="ENVIRONMENT"
    )


def configure_logging(environment: Environment) -> BoundLogger:
    logging.basicConfig(
        format="%(message)s",
        stream=sys.stdout,
        level=environment.log_level,
        force=True,
    )
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
    return structlog.get_logger()


def main():
    load_dotenv()
    default_environment = Environment.DEVELOPMENT
    log = configure_logging(default_environment)

    try:
        config = Config()
    except ValidationError as error:
        log.error("Failed to load configuration", error=error.errors())
        sys.exit(1)

    if config.environment != default_environment:
        log = configure_logging(config.environment)

    log.info(
        "Application started",
        api_key=config.api_key,
        environment=config.environment.value,
    )


if __name__ == "__main__":
    main()
