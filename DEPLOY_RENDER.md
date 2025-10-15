Deploying to Render
====================

This repo includes a sample `render.yaml` to deploy a Go web service to Render.

Quick steps
-----------
1. Push your branch to GitHub (already done).
2. In Render dashboard, create a new service -> "Web Service" -> Connect to GitHub and select `tacsq-platform-api` repo and `main` branch.
3. Render should detect the `render.yaml` and use the provided `buildCommand` and `startCommand`.

Notes
-----
- `render.yaml` builds the binary with `go build -o server ./cmd/api` and runs `./server`.
- Render sets a `PORT` environment variable (e.g. "10000"). `internal/config.LoadConfig()` will prefix a colon automatically so the app listens correctly.
- The `GIN_MODE` env var is set to `release` in `render.yaml`.

Manual service setup (if not using render.yaml)
----------------------------------------------
- Build command: `go build -o server ./cmd/api`
- Start command: `./server`
- Port: Render will set `PORT`; no extra config needed.

Troubleshooting
---------------
- Check the Render logs for build or runtime errors.
- If you need environment variables (DATABASE_URL, etc.) add them in the Render dashboard.
