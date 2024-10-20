# Build stage
FROM golang:1.22.3-alpine3.20 AS builder
WORKDIR /app
# Install necessary dependencies
RUN apk add --no-cache curl tar
# Download and install Golang Migrate
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz -o migrate.linux-amd64.tar.gz && \
    tar xzvf migrate.linux-amd64.tar.gz
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.20 
WORKDIR /app 
# Install PostgreSQL client utilities
RUN apk add --no-cache postgresql-client 
COPY --from=builder /app/main . 
COPY --from=builder /app/migrate ./migrate
COPY /db/migration ./migration
COPY app.env .
COPY entry_point.sh .

EXPOSE 8080
ENTRYPOINT ["/app/entry_point.sh"]
CMD ["/app/main"]