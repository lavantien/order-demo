#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/db/migration -database "$DB_SOURCE" -verbose up

echo "run test and populate data"
cd /app && go test -v --count=1 -cover ./...


echo "start the app"
exec "$@"
