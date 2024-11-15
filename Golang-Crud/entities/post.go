package entities

import(
	"time"
)

type Post struct {
	ID 			int			`gorm:"unique;primaryKey;autoIncrement:true`
	User_id 	int			`json:user_id`
	Title 		string		`json:title`
	Status 		string		`json:status`
	Content 	string		`json:content`
	Created_at 	time.Time	`json:created_at`
	Updated_at 	time.Time	`json:updated_at`
}