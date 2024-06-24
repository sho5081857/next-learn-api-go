# Introduction


# Getting started

## Install Golang

Make sure you have Go 1.22.3 or higher installed.

https://golang.org/doc/install

## Install Dependencies
```bash
go mod download
```

## Environment Variables
```bash
cp .env.example .env
```

## Start DB
```bash
task up
# or...
docker compose up -d
```

## Initial Create Table and Insert Data
```bash
task sql-create
# or...
docker compose exec -T db bash -c 'PGPASSWORD=$POSTGRES_PASSWORD psql -U $POSTGRES_USER -d $POSTGRES_DB' < _tools/first.sql
```

## Start app
```bash
go run .
```
