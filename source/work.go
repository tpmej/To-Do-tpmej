package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

// const directory string = "./todo.db"
var db *sql.DB

func checkForTable() bool {
	var dummy string
	var err error = db.QueryRow("SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%' LIMIT 1").Scan(&dummy)

	if err == sql.ErrNoRows {
		return false
	} else if err != nil {
		return false
	} else {
		return true
	}
}

func createTable() bool {
	var err error
	const query string = `
		CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		desc TEXT,
		completed BOOLEAN NOT NULL DEFAULT 0
		);
	`
	_, err = db.Exec(query)

	if err != nil {
		return false
	} else {
		return true
	}
}
