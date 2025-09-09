# Dockerfile for Bun + Hono application
FROM oven/bun:1.1.13

# Set working directory
WORKDIR /app

# Copy package files
COPY package.json pnpm-lock.yaml ./

# Install dependencies
RUN bun install

# Copy source code
COPY . .

# Expose port (default for Hono/Bun is 3000)
EXPOSE 3000

# Start the server
CMD ["bun", "run", "src/index.ts"]
