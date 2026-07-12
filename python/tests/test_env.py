import pytest
from pydantic import ValidationError

from env import Env, Environment


def test_defaults(monkeypatch: pytest.MonkeyPatch) -> None:
    monkeypatch.delenv("ENVIRONMENT", raising=False)
    monkeypatch.delenv("PORT", raising=False)

    env = Env(_env_file=None)

    assert env.environment is Environment.DEVELOPMENT
    assert env.port == 8080


def test_reads_environment_variables(monkeypatch: pytest.MonkeyPatch) -> None:
    monkeypatch.setenv("ENVIRONMENT", "production")
    monkeypatch.setenv("PORT", "3000")

    env = Env(_env_file=None)

    assert env.environment is Environment.PRODUCTION
    assert env.port == 3000


def test_rejects_unknown_environment(monkeypatch: pytest.MonkeyPatch) -> None:
    monkeypatch.setenv("ENVIRONMENT", "staging")

    with pytest.raises(ValidationError):
        Env(_env_file=None)


def test_rejects_non_numeric_port(monkeypatch: pytest.MonkeyPatch) -> None:
    monkeypatch.delenv("ENVIRONMENT", raising=False)
    monkeypatch.setenv("PORT", "not-a-port")

    with pytest.raises(ValidationError):
        Env(_env_file=None)


def test_rejects_out_of_range_port(monkeypatch: pytest.MonkeyPatch) -> None:
    monkeypatch.delenv("ENVIRONMENT", raising=False)
    monkeypatch.setenv("PORT", "70000")

    with pytest.raises(ValidationError):
        Env(_env_file=None)
