package config

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	// If the db connection is already established, return it
	if db != nil {
		return db
	}

	var err error
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/golang_crud?parseTime=true")
	if err != nil {
		panic(err.Error())
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Database connected successfully!")
	return db
}