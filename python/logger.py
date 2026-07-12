"""Structured JSON logging, with the level derived from the environment."""

import logging
import sys

import structlog

from env import Environment


def new_logger(environment: Environment) -> structlog.typing.FilteringBoundLogger:
    level = logging.DEBUG
    if environment is Environment.PRODUCTION:
        level = logging.INFO

    structlog.configure(
        processors=[
            structlog.contextvars.merge_contextvars,
            structlog.processors.add_log_level,
            structlog.processors.TimeStamper(fmt="iso", utc=True),
            structlog.processors.dict_tracebacks,
            structlog.processors.JSONRenderer(),
        ],
        wrapper_class=structlog.make_filtering_bound_logger(level),
        logger_factory=structlog.PrintLoggerFactory(sys.stdout),
        cache_logger_on_first_use=True,
    )

    return structlog.get_logger().bind(environment=environment)
