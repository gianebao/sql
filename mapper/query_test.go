package mapper_test

import (
	"database/sql"
	"fmt"

	"github.com/gianebao/sql/mapper"
	_ "github.com/mattn/go-sqlite3"
)

func ExampleQuery() {
	var (
		db   *sql.DB
		rows []map[string]interface{}
		err  error
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

	if _, rows, err = mapper.Query(db, "SELECT * FROM user"); len(rows) == 3 {
		fmt.Printf("%s %s %s %s\n", rows[0]["id"], rows[0]["name"], rows[0]["age"], rows[0]["created"])
		fmt.Printf("%s %s %s %s\n", rows[1]["id"], rows[1]["name"], rows[1]["age"], rows[1]["created"])
		fmt.Printf("%s %s %s %s\n", rows[2]["id"], rows[2]["name"], rows[2]["age"], rows[2]["created"])
	} else {
		fmt.Println(err)
	}
	// Output:
	// 1 tim 20 2012-10-01T01:00:01Z
	// 2 joe 25 2012-10-02T02:00:02Z
	// 3 bob 30 2012-10-03T03:00:03Z
}
