# ==========================================
# Build Stage
# ==========================================
FROM node:20-alpine AS builder

RUN apk add --no-cache libc6-compat python3 make g++

WORKDIR /app

COPY package.json package-lock.json* tsconfig.json tsconfig.build.json tsup.config.ts ./
RUN npm ci --ignore-scripts

COPY src ./src
COPY prisma ./prisma

RUN npm run build

# ==========================================
# Production Stage
# ==========================================
FROM node:20-alpine AS runner

RUN addgroup --system --gid 1001 nodejs && \
    adduser --system --uid 1001 nodeuser

WORKDIR /app

ENV NODE_ENV=production

COPY --from=builder --chown=nodeuser:nodejs /app/node_modules ./node_modules
COPY --from=builder --chown=nodeuser:nodejs /app/dist ./dist
COPY --from=builder --chown=nodeuser:nodejs /app/package.json ./package.json
COPY --from=builder --chown=nodeuser:nodejs /app/prisma ./prisma

RUN npx tsc --project tsconfig.build.json --declaration --outDir dist

USER nodeuser

EXPOSE 3000

HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD ["wget", "-q", "--spider", "http://localhost:3000/health", "||", "exit", "1"]

CMD ["node", "dist/index.js"]
