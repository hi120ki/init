import pino from "pino";
import * as dotenv from "dotenv";
import { z } from "zod";

type Config = {
  apiKey: string;
};

const EnvSchema = z.object({
  API_KEY: z.string().min(1, "API_KEY is required"),
});

function createLogger() {
  const environment = process.env.ENVIRONMENT;
  const logLevel = environment === "production" ? "info" : "debug";

  return pino({
    level: logLevel,
    base: { name: "app" },
    timestamp: pino.stdTimeFunctions.isoTime,
  });
}

function loadEnv(logger: pino.Logger) {
  const envResult = dotenv.config({ quiet: true });
  if (envResult.error) {
    logger.warn(
      { error: envResult.error.message },
      "Failed to load .env file, using system environment variables",
    );
  } else {
    logger.info("Loaded .env file");
  }
}

function validateEnv(logger: pino.Logger): Config {
  const parsed = EnvSchema.safeParse(process.env);

  if (!parsed.success) {
    const issues = parsed.error.issues.map((i) => ({
      path: i.path.join("."),
      message: i.message,
    }));
    logger.error(
      { error: { issues } },
      "Failed to process environment variables",
    );
    process.exit(1);
  }

  return {
    apiKey: parsed.data.API_KEY,
  };
}

function main() {
  const logger = createLogger();
  loadEnv(logger);
  const config = validateEnv(logger);

  logger.info({ api_key: config.apiKey }, "Application started");
}

main();
