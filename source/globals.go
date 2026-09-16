package main

import "database/sql"

const (
	directory string = "./todo.db" // initDatabase.go, sql.go
)

var (
	db *sql.DB // sql.go
)
