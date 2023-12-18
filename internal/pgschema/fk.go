package pgschema

type ForeignKey struct {
	Name         string `db:"conname"`
	ForeignTable string `db:"confrelid"`
	Columns      []int  `db:"conkey"`
	Foreign      []int  `db:"confkey"`
}
