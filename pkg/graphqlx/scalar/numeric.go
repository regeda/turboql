package scalar

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/jackc/pgx/v5/pgtype"
)

func serializeNumeric(value any) any {
	switch v := value.(type) {
	case pgtype.Numeric:
		return v
	}
	return nil
}

func parseNumeric(value any) any {
	panic("not yet implemented")
}

func parseLiteralNumeric(valueAST ast.Value) any {
	panic("not yet implemented")
}

var Numeric = graphql.NewScalar(graphql.ScalarConfig{
	Name:         "Numeric",
	Description:  "The `Numeric` type represents a decimal number.",
	Serialize:    serializeNumeric,
	ParseValue:   parseNumeric,
	ParseLiteral: parseLiteralNumeric,
})
