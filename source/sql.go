package main

import (
	"fmt"

	_ "modernc.org/sqlite"
)

// ################################################################ Initialisation block.

// Check for table meeting the requirements.
func checkForTable() bool {
	var (
		dummy string
		err   error
	)
	const query string = "SELECT name FROM sqlite_master WHERE type='table' AND name = 'To-Do tpmej' LIMIT 1"

	err = db.QueryRow(query).Scan(&dummy)

	if err != nil {
		return false
	}
	return true
}

// Create the table.
func createTable() bool {
	var err error
	const query string = `
		CREATE TABLE IF NOT EXISTS "To-Do tpmej" (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		desc TEXT,
		completed BOOLEAN NOT NULL DEFAULT 0
		);
	`
	_, err = db.Exec(query)

	if err != nil {
		return false
	}
	return true
}

// Initialise the table. Used in main().
func initTable() bool {
	if checkForTable() {
		return true
	}
	return createTable()
}

// ################################################################ Commands black.

func help() {
	fmt.Println(`
		Help page:

		Basics:
		"help" ∧ "h" ∧ "-help" ∧ "-h" ∧ "--help"          - Prints this page.
		"version" ∧ "v" ∧ "-version" ∧ "-v" ∧ "--version" - Prints version and repository URL.

		Table manipulation:
		"add" ∧ "-a"      [Title] [Description] - Appends a task to the Table.
		"del" ∧ "-d"      [Title]               - Purges a task from the table.
		"complete" ∧ "-c" [Title]               - Updates task's status to "complete".
	`)
}

func version()

func add()

func del()

func complete()
