import express from "express";
import cors from "cors";
import helmet from "helmet";
import morgan from "morgan";
import routes from "@/routes/index.js";
import { errorHandler } from "@/middleware/errorHandler.js";
import { notFoundHandler } from "@/middleware/notFound.js";
import { requestId } from "@/middleware/requestId.js";
import { corsOptions } from "@/middleware/cors.js";
import { rateLimiter } from "@/middleware/rateLimiter.js";
import { getEnv } from "@/config/index.js";
import { initRedis, closeRedis } from "@/clients/redis.client.js";
import { closePostgresPool } from "@/clients/postgres.client.js";

const app = express();

app.use(helmet());
app.use(cors(corsOptions));
app.use(express.json());
app.use(express.urlencoded({ extended: true }));
app.use(requestId);
app.use(morgan("combined"));
app.use(rateLimiter);

app.use("/api/v1", routes);

app.use(notFoundHandler);
app.use(errorHandler);

const port = getEnv().PORT;

async function bootstrap(): Promise<void> {
  await initRedis();
  app.listen(port, () => {
    console.log(`Server running on port ${port}`);
  });
}

bootstrap();

process.on("SIGTERM", async () => {
  await closeRedis();
  await closePostgresPool();
  process.exit(0);
});