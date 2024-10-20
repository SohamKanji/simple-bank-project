#!/bin/sh

# Wait until PostgreSQL is ready
echo "Waiting for PostgreSQL to be ready..."
until pg_isready -h db -p 5432 -U root; do
  sleep 1
done

# Run migrations using Golang Migrate
source /app/app.env
echo "Running database migrations..."
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

# Start the Go server
echo "Starting the Go server..."
exec "$@"
