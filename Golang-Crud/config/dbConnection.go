package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	dsn := os.Getenv("DB_CREDENTIAL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to Database:", err)
		return
	}
	fmt.Println(dsn)
	fmt.Println("Database connection established")
}
