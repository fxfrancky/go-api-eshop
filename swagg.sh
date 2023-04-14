#!/bin/sh
set -e
echo "updating the api swagger doc"

# swag init --dir ./cmd/http,./internal/handlers,./internal/models
swag init --parseDependency --parseInternal --dir ./cmd/http,./internal/handlers,./internal/models

echo "api swagger doc updated"
exec "$@"
