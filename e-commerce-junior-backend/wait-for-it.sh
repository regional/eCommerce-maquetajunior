#!/bin/sh

TIMEOUT=60
HOST="db"
PORT="3306"

echo "Waiting for $HOST:$PORT..."

start_ts=$(date +%s)
while :
do
    echo "Checking $HOST:$PORT..."
    if nc -z "$HOST" "$PORT"; then
        end_ts=$(date +%s)
        echo "$HOST:$PORT is available after $((end_ts - start_ts)) seconds"
        break
    fi
    sleep 1
    current_ts=$(date +%s)
    elapsed_time=$((current_ts - start_ts))
    if [ $elapsed_time -ge $TIMEOUT ]; then
        echo "Timeout occurred after waiting $TIMEOUT seconds for $HOST:$PORT"
        exit 1
    fi
done

exec "$@"