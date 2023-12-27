package pgschema

import (
	"bytes"
	"strconv"

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

func (t Table) InputVar() string {
	return t.Var() + "Input"
}

func (t Table) GoType() string {
	return t.Title()
}

func (t Table) References(schema Schema) []Reference {
	return schema.References[t.Name]
}

func (t Table) SelectSQL(ref ...string) string {
	b := new(bytes.Buffer)
	b.WriteString("select ")
	t.writeColumns(b)
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

func (t Table) InsertSQL() string {
	b := new(bytes.Buffer)
	b.WriteString("insert into ")
	b.WriteString(t.Name)
	b.WriteByte('(')
	t.writeColumns(b)
	b.WriteString(")values(")
	for i := 1; i <= len(t.Columns); i++ {
		if i > 1 {
			b.WriteByte(',')
		}
		b.WriteByte('$')
		b.WriteString(strconv.Itoa(i))
	}
	b.WriteString(")returning ")
	t.writeColumns(b)
	return b.String()
}

func (t Table) DeleteSQL() string {
	b := new(bytes.Buffer)
	b.WriteString("delete from ")
	b.WriteString(t.Name)
	b.WriteString(" where 1=1")
	return b.String()
}

func (t Table) ColumnsSQL() string {
	b := new(bytes.Buffer)
	t.writeColumns(b)
	return b.String()
}

func (t Table) writeColumns(b *bytes.Buffer) {
	for i, c := range t.Columns {
		if i > 0 {
			b.WriteByte(',')
		}
		b.WriteString(c.Name)
	}
}

func (t Table) GraphqlFilterArgs() []graphqlx.Arg {
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

func (t Table) GraphqlColumnArgs() []graphqlx.Arg {
	var args []graphqlx.Arg
	for _, c := range t.Columns {
		args = append(args, graphqlx.Arg{
			Name:    c.Name,
			Type:    c.GraphqlType(),
			NonNull: c.NotNull,
		})
	}
	return args
}
