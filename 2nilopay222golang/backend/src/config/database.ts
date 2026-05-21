import { PrismaClient } from '@prisma/client';

const prismaClient = new PrismaClient({
  log: process.env.NODE_ENV === 'development' ? ['query', 'error', 'warn'] : ['error'],
});

export const db = prismaClient;

// Handle Prisma client shutdown
process.on('beforeExit', async () => {
  await prismaClient.$disconnect();
});

export default db;
