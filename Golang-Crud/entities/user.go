package entities

import (
	"time"
)

type User struct {
	Id 			int			`gorm:"unique;primaryKey;autoIncrement:true"`
	Name 		string		`json:"name"`
	Password 	string		`json:"Password"`
	Email 		string		`json:"Email"`
	Created_at 	time.Time	`json:"Created_at"`
}