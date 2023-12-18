package pgschema

var goTypes = map[string]string{
	"uuid":                        "string",
	"integer":                     "int",
	"timestamp without time zone": "time.Time",
	"timestamp with time zone":    "time.Time",
	"date":                        "time.Time",
	"text":                        "string",
	"boolean":                     "bool",
	"bigint":                      "int64",
	"uuid[]":                      "[]string",
	"character":                   "string",
	"character varying":           "string",
	"bytea":                       "[]byte",
	"numeric":                     "pgtype.Numeric",
}

var graphqlTypes = map[string]string{
	"uuid":                        "graphql.String",
	"integer":                     "graphql.Int",
	"timestamp without time zone": "graphql.DateTime",
	"timestamp with time zone":    "graphql.DateTime",
	"date":                        "scalar.Date",
	"text":                        "graphql.String",
	"boolean":                     "graphql.Boolean",
	"bigint":                      "graphql.Int",
	"uuid[]":                      "graphql.NewList(graphql.String)",
	"character":                   "graphql.String",
	"character varying":           "graphql.String",
	"bytea":                       "graphql.String",
	"numeric":                     "scalar.Numeric",
}
