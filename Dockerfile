# ==== Frontend Build Stage ====
FROM node:20 AS frontend-build

WORKDIR /frontend

# Copy frontend source code
COPY frontend/package.json frontend/pnpm-lock.yaml ./

# Fix a bug with corepack by installing corepack globally
RUN npm i -g corepack@latest

# Install dependencies
RUN corepack enable pnpm && pnpm i --frozen-lockfile

# Copy the rest of the frontend files
COPY frontend/ ./

# Build the frontend
RUN pnpm run build

# ==== Backend Build Stage ====
FROM golang:1.24.1-alpine3.21 AS backend-build

RUN apk add --no-cache make

WORKDIR /backend

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY backend/ ./

COPY --from=frontend-build /frontend/dist static/

RUN ENV=production && make build

# === Runtime Stage ====
FROM alpine:3.21 AS backend-runtime

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

COPY --from=backend-build /backend/server .

EXPOSE 8080

ENV ENV=production
CMD ["./server"]



