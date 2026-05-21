/**
 * Retry an async function with exponential back-off.
 */
export class RetryError extends Error {
  public attempts: number;
  constructor(message: string, attempts: number) {
    super(message);
    this.attempts = attempts;
    this.name    = "RetryError";
  }
}

export interface RetryOptions {
  maxAttempts: number;
  baseDelayMs: number;
  maxDelayMs: number;
}

const DEFAULT_OPTIONS: RetryOptions = {
  maxAttempts: 3,
  baseDelayMs: 1000,
  maxDelayMs:  10000,
};

/**
 * Execute `fn` and retry up to `maxAttempts` times with
 * exponential back-off capped at `maxDelayMs`.
 */
export async function retryAsync<T>(
  fn: () => Promise<T>,
  opts: Partial<RetryOptions> = {},
): Promise<T> {
  const { maxAttempts, baseDelayMs, maxDelayMs } = {
    ...DEFAULT_OPTIONS,
    ...opts,
  };

  let lastError: unknown;

  for (let attempt = 1; attempt <= maxAttempts; attempt += 1) {
    try {
      return await fn();
    } catch (err) {
      lastError = err;
      if (attempt === maxAttempts) break;
      const delay = Math.min(baseDelayMs * 2 ** (attempt - 1), maxDelayMs);
      await new Promise((r) => setTimeout(r, delay));
    }
  }

  throw new RetryError(
    `All ${maxAttempts} attempts failed`,
    maxAttempts,
  ).constructor.prototype, { cause: lastError } as Error;
}
