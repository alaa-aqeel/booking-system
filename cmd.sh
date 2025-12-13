#!/bin/sh

# ===============================
# Load .env if exists
# ===============================
if [ -f .env ]; then
  export $(grep -v '^#' .env | xargs)
fi

# ===============================
# Configuration
# ===============================
DB_URL="${DATABASE_URL}"
MIGRATIONS_DIR="database/migrations"

# ===============================
# Validation
# ===============================
if [ -z "$DB_URL" ]; then
  echo "❌ DATABASE_URL is not set"
  exit 1
fi

# ===============================
# Commands
# ===============================
case "$1" in
  migrate:up)
    echo "🚀 Running migrations UP..."
    migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" up
    ;;

  migrate:down)
    echo "⬇️  Rolling back last migration..."
    migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" down 1
    ;;

  migrate:reset)
    echo "🔥 Resetting database..."
    migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" down
    migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" up
    ;;

  version)
    migrate -path "$MIGRATIONS_DIR" -database "$DB_URL" version
    ;;

  migrate:create)
    if [ -z "$2" ]; then
      echo "❌ Migration name required: ./migrate.sh create migration_name"
      exit 1
    fi
    echo "✏️ Creating new migration: $2"
    migrate create -ext sql -dir "$MIGRATIONS_DIR" "$2"
    ;;

  *)
    echo "Usage: ./migrate.sh [up|down|reset|version|create migration_name]"
    exit 1
    ;;
esac

echo "✅ Done"
