package databases

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"os"
)

var (
	DB  *sql.DB
	err error
)

func ConnectDB() {
	fmt.Println("Connecting to database.....")

	// set dn for connecting into mysql
	dsn := os.Getenv("MYSQL_URL")
	if dsn == "" {
		fmt.Println("MYQL_URL env variable wasn't loaded/set")
	}

	// connect into mysql database driver
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Error opening database: %s", err.Error()))
	}

    err := DB.Ping(); 
	if err != nil {
		panic(fmt.Sprintf("Error connecting to the database: %s", err.Error()))
	}

    // defer the close till after the main function has finished
	fmt.Println("Successfully connected to the database")
}