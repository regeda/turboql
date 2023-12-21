# TurboQL

TurboQL is a tool designed to automatically generate GraphQL schemas tailored to your PostgreSQL database. It exclusively supports PostgreSQL databases.

The generator seamlessly translates the information schema of your database into corresponding GraphQL models and objects. Notably, TurboQL focuses on creating GraphQL schemas specifically for read-only operations and does not accommodate mutations.

It is crucial to ensure that your database schema is thoroughly designed with valid foreign key constraints. This is essential for establishing meaningful relationships between GraphQL objects during the schema generation process.

## Getting started

Setup the generator:
```
go install github.com/regeda/turboql/cmd/turboqlgen
```

Connect to your database and generate the GraphQL schema:
```
PG_URI='postgres://localhost:5432/postgres?sslmode=disable' turboqlgen
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
