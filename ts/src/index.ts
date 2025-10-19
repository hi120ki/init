import pino, { type Logger } from "pino";
import * as dotenv from "dotenv";
import { z } from "zod";

enum Environment {
  Development = "development",
  Production = "production",
}

const EnvSchema = z.object({
  API_KEY: z.string().min(1, "API_KEY is required"),
  ENVIRONMENT: z.enum(Environment).default(Environment.Development),
});

type Config = {
  apiKey: string;
  environment: Environment;
};

const createLogger = (environment: Environment): Logger => {
  const level = environment === Environment.Production ? "info" : "debug";
  return pino({
    level,
    base: { name: "app" },
    timestamp: pino.stdTimeFunctions.isoTime,
  });
};

const parseConfig = (logger: Logger): Config => {
  const parsed = EnvSchema.safeParse(process.env);
  if (!parsed.success) {
    const issues = parsed.error.issues.map((issue) => ({
      path: issue.path.join(".") || "<root>",
      message: issue.message,
    }));
    logger.error(
      { error: { issues } },
      "Failed to process environment variables",
    );
    process.exit(1);
  }
  return {
    apiKey: parsed.data.API_KEY,
    environment: parsed.data.ENVIRONMENT,
  };
};

const main = (): void => {
  dotenv.config({ quiet: true });
  const initialEnvironment = Environment.Development;
  const bootstrapLogger = createLogger(initialEnvironment);
  const config = parseConfig(bootstrapLogger);
  const logger =
    config.environment === initialEnvironment
      ? bootstrapLogger
      : createLogger(config.environment);

  logger.info(
    { api_key: config.apiKey, environment: config.environment },
    "Application started",
  );
};

main();
