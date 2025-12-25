#!/bin/sh
set -e

echo "Running migrations..."

# Ensure strictly required variables are present
if [ -z "$DB_HOST" ]; then
    echo "Error: DB_HOST is not set"
    exit 1
fi

# Construct DB_URL string for the migrate tool
# Note: Railway provides PG* variables, but we mapped them to DB_* in the Service settings.
export DB_URL="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=$DB_SSLMODE"

# Debug: Print DB Host (partially masked)
echo "----------------------------------------"
echo "STARTING BACKEND SERVICE"
echo "DB_HOST: $DB_HOST"
echo "DB_PORT: $DB_PORT"
echo "DB_USER: $DB_USER"
echo "PORT Variable: $PORT"
echo "----------------------------------------"

# Run migrations (Soft fail - try to start app even if DB is unreachable temporarily)
echo "Running migrations..."
/usr/local/bin/migrate -path /migrations -database "$DB_URL" up
EXIT_CODE=$?

if [ $EXIT_CODE -ne 0 ]; then
    echo "⚠️ MIGRATION FAILED (Exit Code: $EXIT_CODE)"
    echo "Continuing to start application anyway to keep container alive for debugging..."
else
    echo "✅ Migrations completed successfully."
fi

echo "Starting application binary..."

# Run the Go binary and capture output
# We exec it so it handles signals correctly
exec /api
