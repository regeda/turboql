package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/regeda/turboql/examples/bookstore/pkg/bookstore"
)

var (
	listen  = flag.String("listen", "127.0.0.1:8080", "Listen address")
	gqlpath = flag.String("gqlpath", "/graphql", "GraphQL path")
)

func main() {
	flag.Parse()

	cfg, err := pgxpool.ParseConfig(os.Getenv("PG_URI"))
	if err != nil {
		log.Fatalf("Could not parse PG_URI env var: %v", err)
	}

	ctx := context.Background()

	db, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: bookstore.CreateFields(db),
	}
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle(*gqlpath, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(rw, r.WithContext(ctx))
	}))

	log.Printf("Open http://%s%s in your browser", *listen, *gqlpath)

	log.Fatal(http.ListenAndServe(*listen, nil))
}
