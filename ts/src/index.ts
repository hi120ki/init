import { type Env, loadEnv } from "./env.js";
import { createLogger } from "./logger.js";

function main(): void {
  let env: Env;
  try {
    env = loadEnv();
  } catch (error) {
    console.error(
      "Failed to load environment variables:",
      error instanceof Error ? error.message : error,
    );
    process.exit(1);
  }

  const logger = createLogger(env.environment);

  logger.info({ port: env.port }, "application started");
}

main();
