package pgschema

import (
	"bytes"
	"strings"
)

type Table struct {
	Schema      string `db:"schemaname"`
	Name        string `db:"tablename"`
	Columns     []Column
	ForeignKeys []ForeignKey
}

func (t Table) Title() string {
	return strings.Title(t.Name)
}

func (t Table) ColumnAt(i int) (Column, bool) {
	for _, c := range t.Columns {
		if c.Num == i {
			return c, true
		}
	}
	return Column{}, false
}

func (t Table) GraphqlVar() string {
	return t.Name + "Type"
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
	if len(ref) > 0 {
		b.WriteString(" where ")
		b.WriteString(ref[0])
		b.WriteString(" = any($1)")
	}
	return b.String()
}
