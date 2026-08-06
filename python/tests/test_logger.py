import json

import pytest
from env import Environment
from logger import new_logger


def last_log_line(capsys: pytest.CaptureFixture[str]) -> dict:
    lines = capsys.readouterr().out.strip().splitlines()
    return json.loads(lines[-1])


def test_development_logs_debug(capsys: pytest.CaptureFixture[str]) -> None:
    logger = new_logger(Environment.DEVELOPMENT)

    logger.debug("debug message")

    entry = last_log_line(capsys)
    assert entry["event"] == "debug message"
    assert entry["level"] == "debug"
    assert entry["environment"] == "development"


def test_production_suppresses_debug(capsys: pytest.CaptureFixture[str]) -> None:
    logger = new_logger(Environment.PRODUCTION)

    logger.debug("debug message")

    assert capsys.readouterr().out == ""


def test_production_logs_info(capsys: pytest.CaptureFixture[str]) -> None:
    logger = new_logger(Environment.PRODUCTION)

    logger.info("application started", port=8080)

    entry = last_log_line(capsys)
    assert entry["event"] == "application started"
    assert entry["level"] == "info"
    assert entry["environment"] == "production"
    assert entry["port"] == 8080
