package scalar

import (
	"encoding/json"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

type date struct {
	time.Time
}

func (v date) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Time.Format(time.DateOnly))
}

func serializeDate(value any) any {
	switch v := value.(type) {
	case time.Time:
		return date{Time: v}
	}
	return nil
}

func parseDate(value any) any {
	panic("not yet implemented")
}

func parseLiteralDate(valueAST ast.Value) any {
	panic("not yet implemented")
}

var Date = graphql.NewScalar(graphql.ScalarConfig{
	Name:         "Date",
	Description:  "The `Date` type represents a date only format like YYYY-MM-DD.",
	Serialize:    serializeDate,
	ParseValue:   parseDate,
	ParseLiteral: parseLiteralDate,
})
