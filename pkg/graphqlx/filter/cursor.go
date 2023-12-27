package filter

import "github.com/graphql-go/graphql"

var (
	String = scalarInputFilter(graphql.String)
	Int    = scalarInputFilter(graphql.Int)
)

func NewArgumentConfig(name string, filter graphql.InputObjectConfigFieldMap) *graphql.ArgumentConfig {
	return &graphql.ArgumentConfig{
		Type: graphql.NewInputObject(graphql.InputObjectConfig{
			Name:   name,
			Fields: filter,
		}),
	}
}

func NewCursorInput(filter *graphql.ArgumentConfig) graphql.FieldConfigArgument {
	return graphql.FieldConfigArgument{
		"filter": filter,
		"limit": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}
}

func scalarInputFilter(in graphql.Input) *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: in.Name() + "Filter",
		Fields: graphql.InputObjectConfigFieldMap{
			"eq": &graphql.InputObjectFieldConfig{
				Type: in,
			},
			"gt": &graphql.InputObjectFieldConfig{
				Type: in,
			},
			"lt": &graphql.InputObjectFieldConfig{
				Type: in,
			},
			"gte": &graphql.InputObjectFieldConfig{
				Type: in,
			},
			"lte": &graphql.InputObjectFieldConfig{
				Type: in,
			},
		},
	})
}
