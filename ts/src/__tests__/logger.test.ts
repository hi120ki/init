import assert from "node:assert/strict";
import { test } from "node:test";

import { Environment } from "../env.js";
import { createLogger } from "../logger.js";

test("development logger enables debug level", () => {
  const logger = createLogger(Environment.Development);

  assert.equal(logger.level, "debug");
});

test("production logger raises the level to info", () => {
  const logger = createLogger(Environment.Production);

  assert.equal(logger.level, "info");
});

test("logger binds the environment", () => {
  const logger = createLogger(Environment.Production);

  assert.equal(logger.bindings().environment, Environment.Production);
});
