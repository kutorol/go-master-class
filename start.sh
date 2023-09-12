#!/bin/sh

set -e

echo "run db migration"

/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"

# берет все параметры переданные в файл и запускает их
exec "$@"