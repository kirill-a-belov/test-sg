#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
cmd1="$2"
cmd2="$3"

until PGPASSWORD="postgres" psql -h "$host" -U "postgres" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
$cmd1
$cmd2
