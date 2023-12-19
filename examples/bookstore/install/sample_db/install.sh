#!/bin/bash

here=$(dirname $0)

cat $here/*.sql | psql -d "$PG_URI"
