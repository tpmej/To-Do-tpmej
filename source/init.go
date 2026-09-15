package main

import (
	"os"
)

// Define the constant of database's directory
const directory string = "./todo.db"

// Check for database.
func checkForDatabase() bool {
	var err error
	_, err = os.Stat(directory)

	if err == nil {
		return true
	} else if err != nil {
		return false
	}
	return false
}

// Touch the database file.
func touchDatabase() bool {
	var database *os.File
	var err error
	database, err = os.Create(directory)

	if err == nil {
		defer database.Close()
		return true
	} else if err != nil {
		return false
	}
	return false
}

// Open the database.
func openDatabase() (*os.File, bool) {
	var database *os.File
	var err error

	database, err = os.OpenFile(directory, os.O_RDWR, 0666)

	if err == nil {
		return database, true
	} else if err != nil {
		return nil, false
	}
	return nil, false
}

// Initialise the database for main function.
func initDatabase() (*os.File, bool) {
	var database *os.File
	var isOpen bool

	// Check if the database does not exist, if not, touch it.
	if checkForDatabase() == false {
		if touchDatabase() == false {
			return nil, false
		}
	}

	database, isOpen = openDatabase()

	if isOpen == false {
		return nil, false
	} else if isOpen == true {
		return database, true
	}
	return nil, false
}
