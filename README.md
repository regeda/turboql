# TurboQL

TurboQL generates the GraphQL schema for your database. TurboQL supports the PostgreSQL database only.

The generator converts the information schema from the database to the corresponding GraphQL models and objects.

Actually, the generator makes the GraphQL schema for read-only purposes and does not support mutations.

The database schema must be properly designed with the valid foreign key constraints in order to generate a proper relationship between GraphQL objects.

## Getting started

Setup the generator:
```
go get github.com/regeda/turboql/cmd/turboqlgen
```

Connect to your database and generate the GraphQL schema:
```
PG_URI=postgres://localhost:5432/postgres?sslmode=disable turboqlgen
```

## Example

Explore the [Bookstore](/examples/bookstore) example.

The following `Makefile` commands will help you to start from a scratch:

Launch the database:
```
make docker-up
```

Migrate the database schema:
```
make migrate-up
```

Generate the GraphQL schema:
```
make generate
```

Run the web application:
```
make run
```
