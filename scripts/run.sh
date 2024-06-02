#!/bin/bash

CONTAINER_NAME=postgres
POSTGRES_USER=postgres
POSTGRES_DB=demo
WORKDIR=/var/lib/postgresql

docker run -dt \
  --name ${CONTAINER_NAME} \
  -p 5432:5432 \
  -e POSTGRES_USER=${POSTGRES_USER} \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=${POSTGRES_DB} \
  postgres:14

docker cp internal/db/messages/schema.sql postgres:${WORKDIR}
sleep 1
docker exec -u ${POSTGRES_USER} -w ${WORKDIR} ${CONTAINER_NAME} bash -c "psql -f schema.sql -d ${POSTGRES_DB}"
docker exec -u ${POSTGRES_USER} -w ${WORKDIR} ${CONTAINER_NAME} bash -c "psql -d ${POSTGRES_DB} -c '\dt'"
