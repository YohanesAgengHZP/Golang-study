package helper

import (
	// "database/sql"
    "fmt"
    // "log"
	"sync-databases-cass-mysql/connection"
    _ "github.com/go-sql-driver/mysql"
	"time"
)

var chunkSize = 20

func dateTime () {
	// Get current date and time
    currentTime := time.Now()

    // Format date and time
    formattedTime := currentTime.Format("2006-01-02 15:04:05")

    // Print formatted date and time
    fmt.Println(formattedTime)
}

func SyncStatusNumber() {
	fmt.Println("Start syncing for number status")
	fmt.Println("===============================")
	dateTime()

	Connection.ConnectionMySql()
	fmt.Println("")
	Connection.ConnectionCassandra()
}

func QueryMySQL() {
	
}