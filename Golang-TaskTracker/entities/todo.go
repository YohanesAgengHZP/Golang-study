package entities

import (
	// "errors"
	// "fmt"
	"time"
)

type Todo struct {
	Title string 			`json:"title"`
	Completed bool 			`json:"completed"`
	CreatedAt time.Time 	`json:"created_at"`
	CompletedAt *time.Time 	`json:"completed_at"`
}

type Todos []Todo
var todos Todos