#!/bin/bash
docker run --rm -p 5432:5432 -e POSTGRES_HOST_AUTH_METHOD=trust -e POSTGRES_PASSWORD=root postgres:15.4
