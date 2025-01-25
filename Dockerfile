FROM oven/bun:1.1.38-alpine AS build

WORKDIR /app

COPY package.json bun.lockb .
RUN bun install

COPY . .
RUN bun run build

EXPOSE 3000

CMD ["bun", "./build/index.js"]
