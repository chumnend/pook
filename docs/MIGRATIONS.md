# Migrations Guide

For database migrations I am using **golang-migrate**(github.com/golang-migrate/migrate/v4). See [the repo](https://github.com/golang-migrate/migrate).

## Installation

```sh
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate /usr/local/bin/
```

## Usage

To create migrations files,

```sh
migrate create -ext sql -dir db/migrations -seq create_users_table
```

To apply migration files,

```sh
migrate -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" -path db/migrations up
```

To rollback,

```sh
migrate -database "postgres://user:password@localhost:5432/mydb?sslmode=disable" -path db/migrations down 1

```
