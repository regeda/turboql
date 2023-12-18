package pgschema_test

import (
	"os"
	"testing"

	"github.com/regeda/turboql/internal/pgschema"
	"github.com/stretchr/testify/require"
)

func Test_Builder_Execute(t *testing.T) {
	b, err := pgschema.NewBuilder(os.Stdout)

	require.NoError(t, err, "pgschema.NewBuilder")

	err = b.Execute(pgschema.BuilderParams{
		Package: pgschema.Package{
			Name: "test",
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
				},
			},
		),
	})

	require.NoError(t, err)
}
