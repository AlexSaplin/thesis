#!/bin/bash

set -xeuo pipefail

docker network create --driver bridge deepmux || true

docker-compose -f docker-compose-db.yml kill
docker-compose -f docker-compose-db.yml down
docker-compose -f docker-compose-db.yml up -d

sleep 10
echo "applying postgres migrations"
migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path services/ardea/migrations up
migrate -database "postgres://postgres:postgres@localhost:5433/postgres?sslmode=disable" -path services/gorilla/migrations up
migrate -database "postgres://postgres:postgres@localhost:5435/postgres?sslmode=disable" -path services/slav/migrations up
migrate -database "postgres://postgres:postgres@localhost:5436/postgres?sslmode=disable" -path services/ibis/migrations up


echo "setting up minio"

mc config host add myminio http://localhost:9000 minioadmin minioadmin
mc mb myminio/models
mc policy set public myminio/models
mc mb myminio/functions
mc policy set public myminio/functions
