package models

import (
	"fmt"
	"golang-crud/config"
	"golang-crud/entities"
)

func GetAll() []entities.User {
	db := config.ConnectDB()
	
	rows, err := db.Query(`SELECT * FROM users`)
	if err != nil {
		fmt.Println("Error when executing Query")
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		// Scan used to swap variable into user variable
		var user entities.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Created_at); err != nil {

			panic(err)
		}
		users = append(users, user)
	}
	return users
}