package batcher

import (
	"context"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/graphql-go/graphql"
	"github.com/stephenafamo/scan"
	"github.com/stephenafamo/scan/pgxscan"
)

type QueryResolver func(graphql.ResolveParams) (string, []any)

func GraphqlOne[V any](pq pgxscan.Queryer, query QueryResolver) graphql.FieldResolveFn {
	mapper := scan.StructMapper[V]()
	return func(p graphql.ResolveParams) (any, error) {
		sql, args := query(p)
		return pgxscan.One(p.Context, pq, mapper, sql, args...)
	}
}

func GraphqlAll[V any](pq pgxscan.Queryer, query QueryResolver) graphql.FieldResolveFn {
	mapper := scan.StructMapper[V]()
	return func(p graphql.ResolveParams) (any, error) {
		sql, args := query(p)
		return pgxscan.All(p.Context, pq, mapper, sql, args...)
	}
}

func NewLoader[K comparable, V any](pq pgxscan.Queryer, indexer func(V) K, query string) *dataloader.Loader[K, V] {
	mapper := scan.StructMapper[V]()
	return dataloader.NewBatchedLoader(func(ctx context.Context, keys []K) []*dataloader.Result[V] {
		data, err := pgxscan.All(ctx, pq, mapper, query, keys)
		if err != nil {
			return errToResult[K, V](keys, err)
		}
		mm := make(map[K]V, len(data))
		for _, v := range data {
			mm[indexer(v)] = v
		}
		return mapToResult(keys, mm)
	})
}

func NewListLoader[K comparable, V any](pq pgxscan.Queryer, indexer func(V) K, query string) *dataloader.Loader[K, []V] {
	mapper := scan.StructMapper[V]()
	return dataloader.NewBatchedLoader(func(ctx context.Context, keys []K) []*dataloader.Result[[]V] {
		data, err := pgxscan.All(ctx, pq, mapper, query, keys)
		if err != nil {
			return errToResult[K, []V](keys, err)
		}
		mm := make(map[K][]V, len(data))
		for _, v := range data {
			id := indexer(v)
			mm[id] = append(mm[id], v)
		}
		return mapToResult(keys, mm)
	})
}

func mapToResult[K comparable, V any](keys []K, m map[K]V) []*dataloader.Result[V] {
	r := make([]*dataloader.Result[V], len(keys))
	for i, k := range keys {
		r[i] = &dataloader.Result[V]{Data: m[k]}
	}
	return r
}

func errToResult[K comparable, V any](keys []K, err error) []*dataloader.Result[V] {
	r := make([]*dataloader.Result[V], len(keys))
	for i := range keys {
		r[i] = &dataloader.Result[V]{Error: err}
	}
	return r
}
