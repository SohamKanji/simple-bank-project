#!/bin/sh

# Wait until PostgreSQL is ready
echo "Waiting for PostgreSQL to be ready..."
until pg_isready -h db -p 5432 -U root; do
  sleep 1
done

# Run migrations using Golang Migrate
echo "Running database migrations..."
/app/migrate -path /app/migration -database "postgresql://root:secret@db:5432/simple_bank?sslmode=disable" -verbose up

# Start the Go server
echo "Starting the Go server..."
exec "$@"
