# Go API Backend

This project is a backend application built using Go (Golang) to provide APIs for user management and health check functionalities.

## Project Structure

```
go-api-backend
├── cmd
│   └── api
│       └── main.go          # Entry point of the application
├── internal
│   ├── api
│   │   ├── handlers
│   │   │   └── health.go    # Health check handler
│   │   ├── routes
│   │   │   └── routes.go     # API routes setup
│   │   └── middleware
│   │       └── auth.go      # Authentication middleware
│   ├── config
│   │   └── config.go        # Configuration settings
│   ├── database
│   │   └── postgres.go      # PostgreSQL database connection
│   ├── models
│   │   └── user.go          # User model
│   └── service
│       └── user_service.go   # User management service
├── pkg
│   └── logger
│       └── logger.go        # Logging functionality
├── api
│   └── openapi.yaml         # OpenAPI specification
├── migrations
│   └── 0001_init.up.sql    # Database schema initialization
├── scripts
│   └── migrate.sh           # Database migration script
├── test
│   └── integration
│       └── user_test.go     # Integration tests for user API
├── .env.example              # Example environment variables
├── .gitignore                # Git ignore file
├── Dockerfile                # Docker image instructions
├── Makefile                  # Build and run tasks
├── go.mod                    # Go module dependencies
└── README.md                 # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd go-api-backend
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables:**
   Copy the `.env.example` to `.env` and fill in the required values.

4. **Run the application:**
   ```bash
   go run cmd/api/main.go
   ```

5. **Run migrations:**
   ```bash
   ./scripts/migrate.sh
   ```

## Usage

- The API provides endpoints for user management and health checks.
- Refer to the `api/openapi.yaml` for detailed API documentation.

## Testing

- Integration tests can be found in the `test/integration/user_test.go` file.
- Run tests using:
  ```bash
  go test ./...
  ```

## License

This project is licensed under the MIT License. See the LICENSE file for details.