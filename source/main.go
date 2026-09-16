package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		database         *os.File
		initDatabaseBool bool
		initTableBool    bool
	)

	database, initDatabaseBool = initDatabase()
	initTableBool = initTable()

	// Check for database.
	if !initDatabaseBool {
		fmt.Println("File init err.")
		return
	}
	fmt.Println("File initialised.")

	// Check for table
	if !initTableBool {
		fmt.Println("Table init err")
		return
	}
	fmt.Println("Table initialised.")

	fmt.Println(database)
}
