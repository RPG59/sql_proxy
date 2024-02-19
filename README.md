# PostgreSQL Proxy

## Hot reload

```sh
nodemon --exec go run main.go --signal SIGTERM
```

## How to make dump

```sh
pg_dumpall --host=1.1.1.1 --database=my-database --globals-only > globals.sql
pg_dump --column-inserts -h 1.1.1.1 -d my-database > ./test_dump.sql

cat globals.sql test_dump.sql > result.sql
```
