# syntax=docker/dockerfile:1

# ---- Build stage ----
FROM golang:1.21-alpine AS build
WORKDIR /src

# Cache dependencies first.
COPY go.mod go.sum ./
RUN go mod download

# Build a static binary (CGO disabled for portability / Alpine musl).
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/server ./cmd/server

# ---- Runtime stage ----
FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=build /out/server /app/server

ENV PORT=8080
EXPOSE 8080
USER nonroot:nonroot

ENTRYPOINT ["/app/server"]
