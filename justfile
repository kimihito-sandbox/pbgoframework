# Development commands

# Run both Vite and Go servers in parallel
dev:
    #!/usr/bin/env bash
    trap 'kill 0' EXIT
    (cd frontend && pnpm dev) &
    go tool air &
    wait

# Run Vite dev server
vite:
    cd frontend && pnpm dev

# Run Go server with air (hot reload)
server:
    go tool air

# Build frontend for production
build-frontend:
    cd frontend && pnpm build

# Build Go binary
build-go:
    go tool templ generate
    go build -o ./app .

# Build everything for production
build: build-frontend build-go

# Install dependencies
install:
    cd frontend && pnpm install
    go mod tidy

# Generate templ files
generate:
    go tool templ generate

# Clean build artifacts
clean:
    rm -rf tmp/ frontend/dist/ app
