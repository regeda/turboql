package batcher

import (
	"context"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/stephenafamo/scan"
	"github.com/stephenafamo/scan/pgxscan"
)

func NewLoader[K comparable, V any](exec pgxscan.Queryer, indexer func(V) K, query string) *dataloader.Loader[K, V] {
	batchFunc := func(ctx context.Context, keys []K) []*dataloader.Result[V] {
		data, err := pgxscan.All(ctx, exec, scan.StructMapper[V](), query, keys)
		if err != nil {
			return errToResult[K, V](keys, err)
		}
		mm := make(map[K]V, len(data))
		for _, v := range data {
			mm[indexer(v)] = v
		}
		return mapToResult(keys, mm)
	}
	return dataloader.NewBatchedLoader(batchFunc)
}

func NewListLoader[K comparable, V any](exec pgxscan.Queryer, indexer func(V) K, query string) *dataloader.Loader[K, []V] {
	batchFunc := func(ctx context.Context, keys []K) []*dataloader.Result[[]V] {
		data, err := pgxscan.All(ctx, exec, scan.StructMapper[V](), query, keys)
		if err != nil {
			return errToResult[K, []V](keys, err)
		}
		mm := make(map[K][]V, len(data))
		for _, v := range data {
			id := indexer(v)
			mm[id] = append(mm[id], v)
		}
		return mapToResult(keys, mm)
	}
	return dataloader.NewBatchedLoader(batchFunc)
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
