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
	var file *os.File
	var err error
	file, err = os.Create(directory)

	if err == nil {
		defer file.Close()
		return true
	} else if err != nil {
		return false
	}
	return false
}

// Open the database.
func openDatabase() (*os.File, bool) {
	var file *os.File
	var err error

	file, err = os.OpenFile(directory, os.O_RDWR, 0666)

	if err == nil {
		return file, true
	} else if err != nil {
		return nil, false
	}
	return nil, false
}

// Initialise the database for main function.
func initDatabase() (*os.File, bool) {
	var file *os.File
	var isOpen bool

	// Check if the database does not exist, if not, touch it.
	if checkForDatabase() == false {
		if touchDatabase() == false {
			return nil, false
		}
	}

	file, isOpen = openDatabase()

	if isOpen == false {
		return nil, false
	} else if isOpen == true {
		return file, true
	}
	return nil, false
}
