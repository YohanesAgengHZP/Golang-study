package main

import (
	"fmt"
	// "sync-databases-cass-mysql/connection"
	"sync-databases-cass-mysql/helper"
)

func main() {
	fmt.Println("starting app..")
	// mysqlConnection.ConnectionMysql()
	helper.SyncStatusNumber()
}