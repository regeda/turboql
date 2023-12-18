package pgschema

type Schema struct {
	Tables     map[string]Table
	References map[string][]Reference
}

func NewSchema(tables []Table) Schema {
	s := Schema{
		Tables:     make(map[string]Table, len(tables)),
		References: make(map[string][]Reference),
	}

	for _, t := range tables {
		s.Tables[t.Name] = t
	}

	for _, t := range tables {
		for _, fk := range t.ForeignKeys {
			foreignCol, _ := s.Tables[fk.ForeignTable].ColumnAt(fk.Foreign[0])
			col, _ := t.ColumnAt(fk.Columns[0])

			s.References[fk.ForeignTable] = append(
				s.References[fk.ForeignTable],
				Reference{
					Name:          fk.Name,
					Table:         t,
					Column:        col,
					ForeignColumn: foreignCol,
				})
		}
	}

	return s
}

type Reference struct {
	Name          string
	Table         Table
	Column        Column
	ForeignColumn Column
}
