package filter

import (
	"bytes"
	"strconv"

	"github.com/graphql-go/graphql"
)

var filterOpToSQL = map[string]string{
	"eq":  "=",
	"gt":  ">",
	"lt":  "<",
	"gte": ">=",
	"lte": "<=",
}

func SQL(base string, args []any, p graphql.ResolveParams) (string, []any) {
	q := bytes.NewBufferString(base)
	seq := len(args) + 1
	if filter, ok := p.Args["filter"].(map[string]any); ok {
		for name, arg := range filter {
			if m, ok := arg.(map[string]any); ok {
				for filterOp, v := range m {
					if sqlOp, ok := filterOpToSQL[filterOp]; ok {
						// sql
						q.WriteString(" and ")
						q.WriteString(name)
						q.WriteString(sqlOp)
						q.WriteByte('$')
						q.WriteString(strconv.Itoa(seq))
						seq++
						// arg
						args = append(args, v)
						// apply only the first matched operator
						break
					}
				}
			}
		}
	}
	if limit, ok := p.Args["limit"].(int); ok {
		q.WriteString(" limit ")
		q.WriteString(strconv.Itoa(limit))
	}
	return q.String(), args
}
