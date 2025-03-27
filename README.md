# Live-Poilot API

A high-performance API service built with Go and Fiber framework.

## Features

- Fast HTTP server using Fiber
- Database access with Ent ORM
- Dependency injection with Wire
- Structured logging with Zap
- Configuration management with Cleanenv
- Graceful shutdown

## Technology Stack

- Go 1.24.1
- Fiber v2
- Ent ORM
- PostgreSQL/SQLite
- Wire
- Zap
- Sonic (fast JSON encoding/decoding)

## Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Build and run:
   ```bash
   go run cmd/server/main.go
   ```

## Configuration

Edit `config.yml`:

```yaml
server:
  addr: 0.0.0.0:8000  # Server address

database:
  driver: sqlite3     # sqlite3 or postgres
  source: ":memory:?_fk=1"  # Database connection string
  migrate: true       # Auto migrate database
```

## API Documentation

Base URL: `http://localhost:8000/api`

### User

- `GET /user/:id` - Get user by ID

## Development

### Make Commands

```bash
# Install development tools (air, wire)
make init

# Generate code and tidy dependencies
make generate

# Run the application
make run

# Run with hot reload (using air)
make air

# Build the application
make build

# Show available commands
make help
```

### Docker

```bash
docker build -t live-poilot .
docker run -p 8000:8000 live-poilot
```

### Database Migrations

To generate new migrations:
```bash
go run -mod=mod entgo.io/ent/cmd/ent generate ./pkg/ent/schema
```

## License

MIT
