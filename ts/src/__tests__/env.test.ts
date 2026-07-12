import assert from "node:assert/strict";
import { test } from "node:test";

import { Environment, loadEnv } from "../env.js";

test("applies defaults when variables are missing", () => {
  const env = loadEnv({});

  assert.deepEqual(env, { environment: Environment.Development, port: 8080 });
});

test("reads ENVIRONMENT and PORT from the source", () => {
  const env = loadEnv({ ENVIRONMENT: "production", PORT: "3000" });

  assert.deepEqual(env, { environment: Environment.Production, port: 3000 });
});

test("rejects an unknown ENVIRONMENT", () => {
  assert.throws(() => loadEnv({ ENVIRONMENT: "staging" }));
});

test("rejects a non-numeric PORT", () => {
  assert.throws(() => loadEnv({ PORT: "not-a-port" }));
});

test("rejects an out-of-range PORT", () => {
  assert.throws(() => loadEnv({ PORT: "70000" }));
});
