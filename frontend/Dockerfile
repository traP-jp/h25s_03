FROM oven/bun:debian AS builder

WORKDIR /app

COPY package.json bun.lock ./

RUN bun install

COPY . .

ARG VITE_API_BASE_URL="/api"
RUN bun run build-only

FROM caddy:latest AS client

COPY ./Caddyfile /etc/caddy/

COPY --from=builder /app/dist /usr/share/caddy

EXPOSE 80

CMD ["caddy", "run", "--config", "/etc/caddy/Caddyfile"]
