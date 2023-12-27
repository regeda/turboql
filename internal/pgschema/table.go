package pgschema

import (
	"bytes"

	"github.com/iancoleman/strcase"

	"github.com/regeda/turboql/pkg/graphqlx"
)

type Table struct {
	Schema      string `db:"schemaname"`
	Name        string `db:"tablename"`
	Columns     []Column
	PrimaryKeys []PrimaryKey
	ForeignKeys []ForeignKey
}

func (t Table) Title() string {
	return strcase.ToCamel(t.Name)
}

func (t Table) ColumnAt(i int) (Column, bool) {
	for _, c := range t.Columns {
		if c.Num == i {
			return c, true
		}
	}
	return Column{}, false
}

func (t Table) Var() string {
	return strcase.ToLowerCamel(t.Name)
}

func (t Table) GraphqlVar() string {
	return t.Var() + "Type"
}

func (t Table) FilterVar() string {
	return t.Var() + "Filter"
}

func (t Table) GoType() string {
	return t.Title()
}

func (t Table) References(schema Schema) []Reference {
	return schema.References[t.Name]
}

func (t Table) SQL(ref ...string) string {
	b := new(bytes.Buffer)
	b.WriteString("select ")
	for i, c := range t.Columns {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(c.Name)
	}
	b.WriteString(" from ")
	b.WriteString(t.Name)
	b.WriteString(" where 1=1")
	if len(ref) > 0 {
		b.WriteString(" and ")
		b.WriteString(ref[0])
		b.WriteString(" = any($1)")
	}
	return b.String()
}

func (t Table) GraphqlArgs() []graphqlx.Arg {
	var args []graphqlx.Arg
	for _, c := range t.Columns {
		if f, ok := c.FilterType(); ok {
			args = append(args, graphqlx.Arg{
				Name: c.Name,
				Type: f,
			})
		}
	}
	return args
}
