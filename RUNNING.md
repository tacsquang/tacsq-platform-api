Running locally without a database
----------------------------------

This project can run without a Postgres database for lightweight local development.

Prerequisites:
- Go 1.18+

Steps:

```powershell
cd <repo-root>
# download dependencies
go mod tidy
# run the server
go run ./cmd/api
```

The server listens on the address defined in your environment via PORT. You can set it like:

```powershell
$env:PORT = ":8080"
go run ./cmd/api
```

Endpoints available:
- GET /health -> returns OK
- GET /users -> list users (in-memory)
- POST /users -> create user (JSON {"username":"..","email":".."})
