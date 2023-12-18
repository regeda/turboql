package pgschema

import (
	"strings"
)

type Column struct {
	Name    string `db:"attname"`
	Type    string `db:"atttypid"`
	Num     int    `db:"attnum"`
	NotNull bool   `db:"attnotnull"`
}

func (c Column) Title() string {
	return strings.Title(c.Name)
}

func (c Column) GoType() string {
	t, ok := goTypes[c.Type]
	if ok {
		return t
	}
	return c.Type
}

func (c Column) GraphqlType() string {
	t, ok := graphqlTypes[c.Type]
	if ok {
		return t
	}
	return c.Type
}
