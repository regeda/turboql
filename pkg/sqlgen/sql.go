package sqlgen

import (
	"bytes"
	"strconv"
)

func Update(table string, set map[string]any) (string, []any) {
	args := make([]any, 0, len(set))

	b := bytes.NewBufferString("update ")
	b.WriteString(table)
	b.WriteString(" set ")
	for k, v := range set {
		// sql
		if len(args) > 0 {
			b.WriteByte(',')
		}
		b.WriteString(k)
		b.WriteString("=$")
		b.WriteString(strconv.Itoa(len(args) + 1))
		// args
		args = append(args, v)
	}

	b.WriteString(" where 1=1")

	return b.String(), args
}

func Returning(sql string, args []any, columns string) (string, []any) {
	b := bytes.NewBufferString(sql)
	b.WriteString(" returning ")
	b.WriteString(columns)
	return b.String(), args
}
