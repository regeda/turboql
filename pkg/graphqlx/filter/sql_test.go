package filter_test

import (
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/require"

	"github.com/regeda/turboql/pkg/graphqlx/filter"
)

func Test_Filter_SQL(t *testing.T) {
	cases := []struct {
		name         string
		base         string
		args         map[string]any
		expectedSQL  string
		expectedArgs []any
	}{
		{
			name:        "no filter and no limit",
			base:        "select",
			expectedSQL: "select",
		},
		{
			name: "single filter",
			base: "select",
			args: map[string]any{
				"filter": map[string]any{
					"foo": map[string]any{
						"eq": 1,
					},
				},
			},
			expectedSQL:  "select and foo=$1",
			expectedArgs: []any{1},
		},
		{
			name: "only limit",
			base: "select",
			args: map[string]any{
				"limit": 100,
			},
			expectedSQL: "select limit 100",
		},
		{
			name: "filter and limit",
			base: "select",
			args: map[string]any{
				"filter": map[string]any{
					"foo": map[string]any{
						"eq": 1,
					},
				},
				"limit": 100,
			},
			expectedSQL:  "select and foo=$1 limit 100",
			expectedArgs: []any{1},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sql, args := filter.SQL(c.base, nil, graphql.ResolveParams{
				Args: c.args,
			})

			require.Equal(t, c.expectedSQL, sql)
			require.Equal(t, c.expectedArgs, args)
		})
	}
}
