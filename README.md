# Fiber Template API

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

### Swagger UI

Interactive API documentation is available at:  
`http://localhost:8000/swagger/index.html`

To generate Swagger docs:

```bash
swag init -g cmd/server/main.go --output docs
```

### API Endpoints

Base URL: `http://localhost:8000/api`

#### User

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

## Commit Message Guidelines

We follow the [Conventional Commits](https://www.conventionalcommits.org/) standard for commit messages to maintain a clear and consistent commit history. This helps in automated changelog generation and semantic versioning.

### Commit Message Format

**Default Format:**

```text
<type>(<optional scope>): <description>

<optional body>

<optional footer>
```

- **Type**: Indicates the kind of change (e.g., `feat`, `fix`, `refactor`, `style`, `test`, `docs`, `build`, `ops`, `chore`).
- **Scope**: Provides additional context (optional, project-specific).
- **Description**: A concise summary of the change in imperative, present tense (e.g., "add feature" not "added feature").
- **Body**: Motivation for the change and contrast with previous behavior (optional).
- **Footer**: Information about breaking changes or issue references (optional).

**Breaking Changes**: Indicated by an `!` before the `:` in the subject line (e.g., `feat(api)!: remove status endpoint`). Must be described in the footer with `BREAKING CHANGE:`.

**Examples**:

- `feat: add email notifications on new direct messages`
- `feat(shopping cart): add the amazing button`
- `fix(shopping-cart): prevent order an empty shopping cart`
- `perf: decrease memory footprint for unique visitors by using HyperLogLog`
- `build: update dependencies`
- `refactor: implement fibonacci number calculation as recursion`
- `style: remove empty line`

For more details and additional formats (Merge, Revert, Initial Commit), refer to the [full guidelines](https://gist.github.com/qoomon/5dfcdf8eec66a051ecd85625518cfd13).

### Docker

```bash
docker build -t live-pilot .
docker run -p 8000:8000 live-pilot
```

### Database Migrations

To generate new migrations:

```bash
go run -mod=mod entgo.io/ent/cmd/ent generate ./pkg/ent/schema
```

## License

MIT
