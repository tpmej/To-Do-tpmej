package main

import (
	"os"
)

// Check for database.
func checkForDatabase() bool {
	var err error
	_, err = os.Stat(directory)

	return err == nil
}

// Touch the database.
func touchDatabase() bool {
	var (
		database *os.File
		err      error
	)
	database, err = os.Create(directory)

	if err != nil {
		return false
	}
	defer database.Close()
	return true
}

// Open the database.
func openDatabase() (*os.File, bool) {
	var (
		database *os.File
		err      error
	)

	// Open the database file in read-write mode with linux rw-rw-rw- permissions.
	database, err = os.OpenFile(directory, os.O_RDWR, 0666)

	if err != nil {
		return nil, false
	}
	return database, true
}

// Initialise the database. Used in main().
func initDatabase() (*os.File, bool) {
	var (
		database *os.File
		isOpen   bool
	)

	// Check if the database does NOT exist, if NOT, touch it. If NOT created return fail.
	if !checkForDatabase() {
		if !touchDatabase() {
			return nil, false
		}
	}

	database, isOpen = openDatabase()

	if !isOpen {
		return nil, false
	}
	return database, true
}
