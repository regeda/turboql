package pgschema

import "github.com/iancoleman/strcase"

type Column struct {
	Name    string `db:"attname"`
	Type    string `db:"atttypid"`
	Num     int    `db:"attnum"`
	NotNull bool   `db:"attnotnull"`
}

func (c Column) Title() string {
	return strcase.ToCamel(c.Name)

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

func (c Column) FilterType() (string, bool) {
	t, ok := filterTypes[c.Type]
	return t, ok
}
