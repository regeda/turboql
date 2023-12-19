package pgschema

import (
	_ "embed"
	"io"
	"text/template"

	"github.com/pkg/errors"
)

//go:embed template.tmpl
var tpl string

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
	Schema  Schema
	Package Package
}
