package pgschema

import (
	"io"
	"text/template"

	"github.com/pkg/errors"
)

const tpl = `
{{ define "graphql-field" }}
"{{ .Column.Name }}": &graphql.Field{
	Type: {{ .Column.GraphqlType }},
	Resolve: func(p graphql.ResolveParams) (any, error) {
		return p.Source.(*{{ .Table.GoType }}).{{ .Column.Title }}, nil
	},
}
{{- end }}

{{ define "column-model" }}
{{ .Column.Title }} {{ .Column.GoType }}
{{- end }}

{{ define "table-model" }}
type {{ .Table.GoType }} struct {
{{- range .Table.Columns }} {{ template "column-model" (args "Column" .) -}} {{ end }}
}
{{- end }}

{{ define "graphql-object" }}
{{ .Table.GraphqlVar }} := graphql.NewObject(graphql.ObjectConfig{
	Name: "{{ .Table.GoType }}",
	Fields: graphql.Fields{
		{{- range .Table.Columns }} {{ template "graphql-field" (args "Column" . "Table" $.Table) }}, {{ end }}
	},
})
{{- end }}

{{ define "graphql-entry" }}
"{{ .Table.Name }}": &graphql.Field{
	Type: graphql.NewList({{ .Table.GraphqlVar }}),
	Resolve: func(p graphql.ResolveParams) (any, error) {
		return pgxscan.All(p.Context, pq, scan.StructMapper[*{{ .Table.GoType }}](), "{{ .Table.SQL }}")
	},
}
{{- end }}

{{ define "graphql-refs" }}
{{ range .References }}
{{ $ref := . }}

{{ $ref.Table.Name }}{{ $.Table.Name }}Loader := batcher.NewLoader[{{ $ref.ForeignColumn.GoType }}, *{{ $.Table.GoType }}](
	pq,
	func(v *{{ $.Table.GoType }}) {{ $ref.ForeignColumn.GoType }} {
		return v.{{ $ref.ForeignColumn.Title }}
	},
	"{{ $.Table.SQL $ref.ForeignColumn.Name }}",
)
{{ $ref.Table.GraphqlVar }}.AddFieldConfig("{{ $.Table.Name }}", &graphql.Field{
	Type: {{ $.Table.GraphqlVar }},
	Resolve: func(p graphql.ResolveParams) (any, error) {
		thunk := {{ $ref.Table.Name }}{{ $.Table.Name }}Loader.Load(p.Context, p.Source.(*{{ $ref.Table.GoType }}).{{ $ref.Column.Title }})
		return func() (any, error) { return thunk() }, nil
	},
})


{{ $.Table.Name }}{{ $ref.Table.Name }}Loader := batcher.NewListLoader[{{ $ref.Column.GoType }}, *{{ $ref.Table.GoType }}](
	pq,
	func(v *{{ $ref.Table.GoType }}) {{ $ref.Column.GoType }} {
		return v.{{ $ref.Column.Title }}
	},
	"{{ $ref.Table.SQL $ref.Column.Name }}",
)
{{ $.Table.GraphqlVar }}.AddFieldConfig("{{ $ref.Name }}", &graphql.Field{
	Type: graphql.NewList({{ $ref.Table.GraphqlVar }}),
	Resolve: func(p graphql.ResolveParams) (any, error) {
		thunk := {{ $.Table.Name }}{{ $ref.Table.Name }}Loader.Load(p.Context, p.Source.(*{{ $.Table.GoType }}).{{ $ref.ForeignColumn.Title }})
		return func() (any, error) { return thunk() }, nil
	},
})
{{ end }}
{{- end }}

{{ define "turboql" }}
package {{ .Package.Name }}

import (
	"github.com/graphql-go/graphql"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/regeda/turboql/pkg/batcher"
	"github.com/regeda/turboql/pkg/graphqlx/scalar"
	"github.com/stephenafamo/scan"
	"github.com/stephenafamo/scan/pgxscan"
)

{{ range .Schema.Tables }} {{ template "table-model" (args "Table" .) }} {{ end }}

func CreateFields(pq pgxscan.Queryer) graphql.Fields {
	{{ range .Schema.Tables }} {{ template "graphql-object" (args "Table" .) }} {{ end }}

	{{ range .Schema.Tables }} {{ template "graphql-refs" (args "Table" . "References" (.References $.Schema)) }} {{ end }}

	return graphql.Fields{
		{{ range .Schema.Tables }} {{ template "graphql-entry" (args "Table" .) }},  {{ end }}
	}
}
{{- end }}
`

var funcMap = template.FuncMap{
	"args": func(v ...any) (any, error) {
		if len(v)%2 != 0 {
			return nil, errors.Errorf("odd number %d of map func, must be even", len(v))
		}
		m := make(map[string]any, len(v)/2)
		for i := 0; i < len(v); i += 2 {
			s, ok := v[i].(string)
			if !ok {
				return nil, errors.Errorf("the key (%d) must be a string, given %T", i, v[i])
			}
			m[s] = v[i+1]
		}
		return m, nil
	},
}

type Builder struct {
	w   io.Writer
	tpl *template.Template
}

func NewBuilder(w io.Writer) (*Builder, error) {
	t, err := template.New("turboql").Funcs(funcMap).Parse(tpl)
	if err != nil {
		return nil, errors.WithMessage(err, "template.New")
	}
	return &Builder{
		w:   w,
		tpl: t,
	}, nil
}

func (b *Builder) Execute(params BuilderParams) error {
	return b.tpl.ExecuteTemplate(b.w, "turboql", params)
}

type BuilderParams struct {
	Package Package
	Schema  Schema
}
