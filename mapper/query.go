package mapper

import (
	"database/sql"
	"fmt"
)

func mapRows(cols []string, rows *sql.Rows) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)
	lenCols := len(cols)

	for j := 0; rows.Next(); j++ {
		rawResult := make([][]byte, lenCols)
		values := make([]interface{}, lenCols)

		for i := range rawResult {
			values[i] = &rawResult[i]
		}

		if err := rows.Scan(values...); err != nil {
			panic(err)
		}

		result = append(result, map[string]interface{}{})
		for i, raw := range rawResult {
			result[j][cols[i]] = raw
		}
	}

	return result
}

// Query executes an SQL query and returns the result
func Query(db *sql.DB, query string, args ...interface{}) (cols Fields, result []map[string]interface{}, err error) {
	var (
		rows *sql.Rows
	)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()

	if rows, err = db.Query(query, args...); err != nil {
		return cols, result, err
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return cols, result, err
	}

	if cols, err = ParseFields(rows); err != nil {
		return cols, result, err
	}

	result = mapRows(cols.Names, rows)

	return cols, result, nil
}
