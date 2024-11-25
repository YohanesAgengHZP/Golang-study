package main

import (
	"fmt"
	"os"
	"tasktracker/controllers"
	"tasktracker/storage"
	"tasktracker/utils"
)

func main() {
	// Initialize Todos and Storage
	todos := controllers.Todos{}
	storage := storage.NewStorage[controllers.Todos]("todos.json")

	// Load existing tasks
	if err := storage.Load(&todos); err != nil {
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}

	// Parse and execute command flags
	cmdFlags := utils.NewCommandFlags()
	if err := cmdFlags.Execute(&todos); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}

	// Save updated tasks
	if err := storage.Save(todos); err != nil {
		fmt.Println("Error saving tasks:", err)
		os.Exit(1)
	}
}
