package pgschema

import (
	"context"

	"github.com/pkg/errors"
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
	pg_catalog.pg_attribute
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
	and contype = 'f'
	and confrelid > 0
`
		pksql = `
select
	conname,
	conkey
from
	pg_catalog.pg_constraint
where
	conrelid = $1::regclass
	and contype = 'p'
`
	)

	var (
		tableMapper  = scan.StructMapper[Table]()
		columnMapper = scan.StructMapper[Column]()
		fkMapper     = scan.StructMapper[ForeignKey]()
		pkMapper     = scan.StructMapper[PrimaryKey]()
	)

	tables, err := pgxscan.All(ctx, pq, tableMapper, tablessql, schema)
	if err != nil {
		return nil, errors.WithMessage(err, "scan tables")
	}

	for i, t := range tables {
		tables[i].Columns, err = pgxscan.All(ctx, pq, columnMapper, columnssql, t.Name)
		if err != nil {
			return nil, errors.WithMessagef(err, "scan columns for %q", t.Name)
		}

		tables[i].ForeignKeys, err = pgxscan.All(ctx, pq, fkMapper, fksql, t.Name)
		if err != nil {
			return nil, errors.WithMessagef(err, "scan foreign keys for %q", t.Name)
		}

		tables[i].PrimaryKeys, err = pgxscan.All(ctx, pq, pkMapper, pksql, t.Name)
		if err != nil {
			return nil, errors.WithMessagef(err, "scan primary keys for %q", t.Name)
		}
	}

	return tables, nil
}
