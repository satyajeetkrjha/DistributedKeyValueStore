#!/bin/bash
set -e

trap 'killall distrikv' SIGINT

cd $(dirname $0)

killall distrikv || true
sleep 0.1

go run main.go -db-location=California.db -http-addr=127.0.0.1:8080 -config-file=sharding.toml -shard=California &
go run main.go -db-location=NewYork.db -http-addr=127.0.0.1:8081 -config-file=sharding.toml -shard=NewYork &
go run main.go -db-location=Washington.db -http-addr=127.0.0.1:8082 -config-file=sharding.toml -shard=Washington &

wait