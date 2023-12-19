package pgschema

import (
	"context"

	"github.com/stephenafamo/scan"
	"github.com/stephenafamo/scan/pgxscan"
)

func Scan(ctx context.Context, pq pgxscan.Queryer, schema string) ([]Table, error) {
	const (
		tablessql = `
select
	schemaname,
	tablename
from
	pg_catalog.pg_tables
where
	schemaname = $1`
		columnssql = `
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
	attnum`
		fksql = `
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
`
	)

	var (
		tableMapper  = scan.StructMapper[Table]()
		columnMapper = scan.StructMapper[Column]()
		fkMapper     = scan.StructMapper[ForeignKey]()
	)

	tables, err := pgxscan.All(ctx, pq, tableMapper, tablessql, schema)
	if err != nil {
		return nil, err
	}

	for i, t := range tables {
		tables[i].Columns, err = pgxscan.All(ctx, pq, columnMapper, columnssql, t.Name)
		if err != nil {
			return nil, err
		}

		tables[i].ForeignKeys, err = pgxscan.All(ctx, pq, fkMapper, fksql, t.Name)
		if err != nil {
			return nil, err
		}
	}

	return tables, nil
}
