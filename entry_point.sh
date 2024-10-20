#!/bin/sh

# Run migrations using Golang Migrate
source /app/app.env
echo "Running database migrations..."
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

# Start the Go server
echo "Starting the Go server..."
exec "$@"
