package pgschema

type PrimaryKey struct {
	Name    string `db:"conname"`
	Columns []int  `db:"conkey"`
}
