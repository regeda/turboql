package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/regeda/turboql/internal/pgschema"
)

var (
	packageName = flag.String("package-name", "turboql", "Go package name of the generated files")
	pgSchema    = flag.String("pg-schema", "public", "The schema name of postgres tables")
)

func main() {
	flag.Parse()

	cfg, err := pgx.ParseConfig(os.Getenv("PG_URI"))
	if err != nil {
		log.Fatalf("Could not parse PG_URI env var: %v", err)
	}

	ctx := context.Background()

	db, err := pgx.ConnectConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v, check your PG_URI environment variable", err)
	}

	tables, err := pgschema.Scan(ctx, db, *pgSchema)
	if err != nil {
		log.Fatalf("Could not scan the schema %q: %v", *pgSchema, err)
	}

	b, err := pgschema.NewBuilder(os.Stdout)
	if err != nil {
		log.Fatalf("Could not create the schema builder: %v", err)
	}

	if err := b.Execute(pgschema.BuilderParams{
		Package: pgschema.Package{
			Name: *packageName,
		},
		Schema: pgschema.NewSchema(tables),
	}); err != nil {
		log.Fatalf("Could not create the schema file: %v", err)
	}
}
