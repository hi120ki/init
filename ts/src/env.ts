import { z } from "zod";

export const Environment = {
  Development: "development",
  Production: "production",
} as const;

export type Environment = (typeof Environment)[keyof typeof Environment];

const envSchema = z.object({
  ENVIRONMENT: z
    .enum([Environment.Development, Environment.Production])
    .default(Environment.Development),
  PORT: z.coerce.number().int().min(1).max(65535).default(8080),
});

export type Env = {
  environment: Environment;
  port: number;
};

export function loadEnv(source: NodeJS.ProcessEnv = process.env): Env {
  try {
    process.loadEnvFile();
  } catch {
    // The .env file is optional.
  }

  const parsed = envSchema.safeParse(source);
  if (!parsed.success) {
    throw new Error(z.prettifyError(parsed.error));
  }

  return {
    environment: parsed.data.ENVIRONMENT,
    port: parsed.data.PORT,
  };
}
