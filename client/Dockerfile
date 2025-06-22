# Build stage
FROM node:22.16.0-alpine AS builder

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./

RUN npm ci

COPY . .
RUN npm run build

# Production stage
FROM caddy:alpine

COPY --from=builder /app/dist /srv

COPY ./docker/Caddyfile /etc/caddy/Caddyfile

EXPOSE 80

CMD ["caddy", "run", "--config", "/etc/caddy/Caddyfile", "--adapter", "caddyfile"]
