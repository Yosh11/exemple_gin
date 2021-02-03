#!/bin/sh
# migrate.sh

echo "Apply database migrations"
make migrate_up

exec $cmd
