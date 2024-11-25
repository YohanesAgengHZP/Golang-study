package controllers

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"tasktracker/entities"
	"time"

	"github.com/aquasecurity/table"
)

type Todos []entities.Todo

//method add tasks 
func (todos *Todos) Add(title string) error{
	todo := entities.Todo{
		Title: title,
		Completed: false,
		CompletedAt: nil,
		CreatedAt: time.Now(),
	}
	//append tasks into todos struct
	*todos = append(*todos, todo)

	fmt.Println("Task added successfully")
	return nil
}

//create helper method validate index
func (todos *Todos) ValidateIndex(index int) error{
	if index < 0 || index >= len(*todos){
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

//todos delete method
func (todos *Todos) Delete(index int) error {
	t := *todos
	
	if err := t.ValidateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	fmt.Println("Task has been deleted !")
	return nil
}

//todos display method
func (todos *Todos) Display() {
	for _, todo := range *todos {
		fmt.Printf("Title: %s\nCompleted: %t\nCreatedAt: %s\nCompletedAt: %v\n\n",
			todo.Title, todo.Completed, todo.CreatedAt.Format(time.RFC1123), todo.CompletedAt)
	}
}

func (todos *Todos) Toggle(index int) error {
	// Validate index
	if err := todos.ValidateIndex(index); err != nil {
		return err
	}
	// Get the todo item by reference
	item := &(*todos)[index]

	// Toggle CompletedAt
	if item.CompletedAt == nil {
		// Mark as completed
		completionTime := time.Now()
		item.CompletedAt = &completionTime
		item.Completed = true
	} else {
		// Mark as incomplete
		item.CompletedAt = nil
		item.Completed = false
	}
	fmt.Println("Task has been completed !")
	return nil
}

func (todos *Todos) Edit(index int, title string, toggleCompletion bool) error {
	// Validate index
	if err := todos.ValidateIndex(index); err != nil {
		return err
	}
	item := &(*todos)[index]
	 // reference into items

	item.Title = title
	if toggleCompletion {
		if item.CompletedAt == nil {
			completionTime := time.Now()
			item.CompletedAt = &completionTime
		} else {
			item.CompletedAt = nil
		}
	}
	fmt.Println("Task has been edited !")
	return nil
}

func (todos *Todos) Print() {
	// Create a new table
	t := table.New(os.Stdout)

	// Configure table settings
	t.SetHeaders("No", "Title", "Status", "Created At", "Completed At")
	t.SetRowLines(false)

	for index, todo := range *todos {
		status := "❌" 
		completedAt := "-"
		if todo.Completed {
			status = "✔"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC3339)
			}
		}
		t.AddRow(strconv.Itoa(index), todo.Title, status, todo.CreatedAt.Format(time.RFC3339), completedAt)
	}

	t.Render()
}
