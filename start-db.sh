#!/bin/bash
docker run --rm --shm-size=1g -p 5433:5432 -e POSTGRES_HOST_AUTH_METHOD=trust -e POSTGRES_PASSWORD=root postgres:16.0
