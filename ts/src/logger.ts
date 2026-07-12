import { pino, type Logger } from "pino";

import { Environment } from "./env.js";

export function createLogger(environment: Environment): Logger {
  const level = environment === Environment.Production ? "info" : "debug";

  return pino({
    level,
    timestamp: pino.stdTimeFunctions.isoTime,
    formatters: {
      level: (label) => ({ level: label }),
    },
    base: null,
  }).child({ environment });
}
