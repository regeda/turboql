package main

import (
	"context"
	"flag"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/regeda/turboql/internal/pgschema"
	"github.com/stephenafamo/scan"
	"github.com/stephenafamo/scan/pgxscan"
)

var (
	packageName = flag.String("package-name", "", "Go package name of the generated files")
	pgSchema    = flag.String("pg-schema", "public", "The schema name of postgres tables")
)

func main() {
	flag.Parse()

	if *packageName == "" {
		panic("Specify package-name")
	}

	if *pgSchema == "" {
		panic("Specify pg-schema")
	}

	cfg, err := pgx.ParseConfig(os.Getenv("PG_URI"))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	db, err := pgx.ConnectConfig(ctx, cfg)
	if err != nil {
		panic(err)
	}

	tables, err := pgxscan.All(ctx, db, scan.StructMapper[pgschema.Table](), "select schemaname, tablename from pg_catalog.pg_tables where schemaname = $1", *pgSchema)
	if err != nil {
		panic(err)
	}

	for i, t := range tables {
		columns, err := pgxscan.All(ctx, db, scan.StructMapper[pgschema.Column](), `
select
	attname,
    atttypid::regtype,
	attnum,
	attnotnull
from
	pg_attribute
where
	attrelid = $1::regclass
	and attnum > 0
	and not attisdropped
order by
	attnum`,
			t.Name)
		if err != nil {
			panic(err)
		}
		tables[i].Columns = columns

		foreignKeys, err := pgxscan.All(ctx, db, scan.StructMapper[pgschema.ForeignKey](), `
select
	conname,
	confrelid::regclass,
	conkey,
	confkey
from
	pg_catalog.pg_constraint
where
	conrelid = $1::regclass
	and confrelid > 0
`,
			t.Name)
		if err != nil {
			panic(err)
		}
		tables[i].ForeignKeys = foreignKeys
	}

	b, err := pgschema.NewBuilder(os.Stdout)
	if err != nil {
		panic(err)
	}

	if err := b.Execute(pgschema.BuilderParams{
		Package: pgschema.Package{
			Name: *packageName,
		},
		Schema: pgschema.NewSchema(tables),
	}); err != nil {
		panic(err)
	}
}
