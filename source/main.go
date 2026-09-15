package main

import (
	"fmt"
)

func main() {
	var database, initSuccess = initDatabase()
	if initSuccess == false {
		fmt.Println("Database init err")
	} else {
		fmt.Println("Database initiated!")
	}

}
