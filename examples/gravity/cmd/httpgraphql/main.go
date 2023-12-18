package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/regeda/turboql/examples/gravity/internal/gravity"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	flag.Parse()

	ctx := context.Background()

	cfg, err := pgxpool.ParseConfig(os.Getenv("PG_URI"))
	if err != nil {
		panic(err)
	}

	db, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		panic(err)
	}

	fields := gravity.CreateFields(db)

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(rw, r.WithContext(ctx))
	}))
	http.ListenAndServe(":8080", nil)
}
