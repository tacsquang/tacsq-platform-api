#!/bin/bash

# This script runs the database migrations using the `migrate` tool.

set -e

# Define the database connection string
DB_URL="postgres://user:password@localhost:5432/dbname?sslmode=disable"

# Run the migrations
migrate -path ./migrations -database "$DB_URL" up

echo "Migrations applied successfully."