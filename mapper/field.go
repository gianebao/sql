package mapper

import "database/sql"

// Fields represents collection of fields
type Fields struct {
	Names []string
	Map   map[string]Field
}

// Field represents a definition of a field
type Field struct {
	Name string
	Type string
}

// ParseFields creates a list of field definition
func ParseFields(rows *sql.Rows) (Fields, error) {
	var (
		colTypes []*sql.ColumnType
		columns  []string
		err      error
	)

	if colTypes, err = rows.ColumnTypes(); err != nil {
		return Fields{}, err
	}

	if columns, err = rows.Columns(); err != nil {
		return Fields{}, err
	}

	c := Fields{
		Names: columns,
		Map:   map[string]Field{},
	}

	for i, v := range c.Names {
		c.Map[v] = Field{
			Name: v,
			Type: colTypes[i].DatabaseTypeName(),
		}
	}

	return c, nil
}
