# TurboQL

TurboQL is a tool designed to automatically generate GraphQL schemas tailored to your PostgreSQL database. It exclusively supports PostgreSQL databases.

The generator seamlessly translates the information schema of your database into corresponding GraphQL models and objects. Notably, TurboQL facilitates the complete implementation of CRUD operations (create/read/update/delete) when generating GraphQL schemas.

It is crucial to ensure that your database schema is thoroughly designed with valid foreign key constraints. This is essential for establishing meaningful relationships between GraphQL objects during the schema generation process.

## Getting started

Setup the generator:
```
go install github.com/regeda/turboql/cmd/turboqlgen
```

> How to install [Golang](https://go.dev/doc/install)?

Connect to your database and generate the GraphQL schema:
```
PG_URI='postgres://localhost:5432/postgres?sslmode=disable' turboqlgen
```

> Run `turboqlgen --help` for help on the documentation.
> Or [Create a new issue](https://github.com/regeda/turboql/issues/new).

## Playground

Clone [our](https://github.com/regeda/turboql) repository:
```
git clone https://github.com/regeda/turboql.git
```

Explore the [Bookstore](/examples/bookstore) example.

The following `Makefile` commands will help you to **start from a scratch**:

### Default Mode

Execute the following command from the repository's root where [`Makefile`](/Makefile) is located:

```
make
```

### Launch the database

Make sure before that your Docker or a virtual machine with the Docker are running!
1. [Rancher Desktop](https://rancherdesktop.io/)
2. [Docker Engine](https://docs.docker.com/engine/install/)

```
make docker-up
```

### Migrate the database schema
```
make migrate-up
```

### Generate the GraphQL schema
```
make generate
```

### Run the web application
```
make run
```

