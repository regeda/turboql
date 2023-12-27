package pgschema_test

import (
	"bytes"
	"fmt"
	"go/format"
	"testing"

	"github.com/regeda/turboql/internal/pgschema"

	"github.com/stretchr/testify/require"
)

func Test_Builder_Execute(t *testing.T) {
	buf := new(bytes.Buffer)

	b, err := pgschema.NewBuilder(buf)

	require.NoError(t, err)

	err = b.Execute(pgschema.BuilderParams{
		Package: pgschema.Package{
			Name: "pgschema_test",
		},
		Schema: pgschema.NewSchema(
			[]pgschema.Table{
				{
					Name: "foobar",
					Columns: []pgschema.Column{
						{
							Name: "foo",
							Type: "integer",
							Num:  1,
						},
						{
							Name: "bar",
							Type: "text",
							Num:  2,
						},
						{
							Name: "baz_ref",
							Type: "integer",
							Num:  3,
						},
					},
					PrimaryKeys: []pgschema.PrimaryKey{
						{
							Name:    "pk_foobar",
							Columns: []int{1, 2},
						},
					},
					ForeignKeys: []pgschema.ForeignKey{
						{
							Name:         "bazquux_fk",
							ForeignTable: "bazquux",
							Foreign:      []int{1},
							Columns:      []int{3},
						},
					},
				},
				{
					Name: "bazquux",
					Columns: []pgschema.Column{
						{
							Name: "baz",
							Type: "integer",
							Num:  1,
						},
						{
							Name: "quux",
							Type: "text",
							Num:  2,
						},
					},
					PrimaryKeys: []pgschema.PrimaryKey{
						{
							Name:    "pk_bazquux",
							Columns: []int{1},
						},
					},
				},
			},
		),
	})

	require.NoError(t, err)

	formattedContent, err := format.Source(buf.Bytes())

	require.NoError(t, err)

	fmt.Println(string(formattedContent))
}
