package mapper_test

import (
	"database/sql"
	"fmt"

	"github.com/gianebao/sql/mapper"
	_ "github.com/mattn/go-sqlite3"
)

func ExampleParseFields() {
	var (
		db     *sql.DB
		rows   *sql.Rows
		err    error
		fields mapper.Fields
	)

	if db, err = sql.Open("sqlite3", ":memory:"); err != nil {
		panic(err)
	}

	db.Exec(`
    CREATE TABLE user (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name VARCHAR(64) NULL,
        age INTEGER NULL,
        created DATETIME NULL
    )`)

	db.Exec(`INSERT INTO user(name, age, created)
    VALUES
    (?, ?, ?), (?, ?, ?), (?, ?, ?)`,
		"tim", "20", "2012-10-01 01:00:01",
		"joe", "25", "2012-10-02 02:00:02",
		"bob", "30", "2012-10-03 03:00:03")

	if rows, err = db.Query("SELECT * FROM user"); err != nil {
		panic(err)
	}

	defer rows.Close()

	if fields, err = mapper.ParseFields(rows); err == nil {
		fmt.Println(fields.Map["id"].Name, fields.Map["id"].Type)
		fmt.Println(fields.Map["name"].Name, fields.Map["name"].Type)
		fmt.Println(fields.Map["age"].Name, fields.Map["age"].Type)
		fmt.Println(fields.Map["created"].Name, fields.Map["created"].Type)
	} else {
		fmt.Println(err)
	}
	// Output:
	// id INTEGER
	// name VARCHAR(64)
	// age INTEGER
	// created DATETIME
}
