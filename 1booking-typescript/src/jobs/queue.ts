import { Queue } from "bullmq";
import { getRedisClient } from "@/clients/redis.client.js";

const connection = getRedisClient();
export const emailQueue   = new Queue("email",    { connection });
export const notifyQueue  = new Queue("notify",   { connection });
export const bookingQueue = new Queue("bookings", { connection });

export default { emailQueue, notifyQueue, bookingQueue };