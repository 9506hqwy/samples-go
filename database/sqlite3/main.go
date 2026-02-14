package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DataMin = 0
	DataMax = 100
)

func main() {
	db, err := sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = createTable(db)
	if err != nil {
		log.Fatal(err)
	}

	err = querySelectAll(db)
	if err != nil {
		log.Fatal(err)
	}

	err = querySelectCond(db)
	if err != nil {
		log.Fatal(err)
	}
}

func createTable(db *sql.DB) error {
	// Create table.
	sqlStmt := `
	CREATE TABLE test (
		id INTERGER NOT NULL PRIMARY KEY,
		name TEXT
	);
	`

	_, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}

	// Insert row into table.
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO test(id, name) values(?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	for i := DataMin; i < DataMax; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("Hello, World %03d", i))
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func querySelectAll(db *sql.DB) error {
	rows, err := db.Query("SELECT id, name FROM test")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string

		err = rows.Scan(&id, &name)
		if err != nil {
			return err
		}

		_, err = fmt.Println(id, name)
		if err != nil {
			return err
		}
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	return nil
}

func querySelectCond(db *sql.DB) error {
	stmt, err := db.Prepare("SELECT name FROM test WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		return err
	}

	_, err = fmt.Println(name)
	if err != nil {
		return err
	}

	return nil
}
