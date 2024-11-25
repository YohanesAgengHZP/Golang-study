package main

import (
	"context"
	"fmt"
	"generic_crud/database"
	"generic_crud/initializer"
	"generic_crud/router" // Import your router package
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	wd, err := os.Getwd()
	//c check for working current directory .env
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}
	log.Printf("Current working directory: %s", wd)

	initializer.LoadEnvVar()
	databases.ConnectDB()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Initialize the router with all routes
	r := router.InitRouter()

	// Create the HTTP server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Coba implementasi channel for goroutines
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start the server
	go func() {
		fmt.Println("Application running...")
		fmt.Printf("App is running on port %s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v\n", err)
		}
	}()

	// sigterm
	<-stop
	fmt.Println("Shutting down...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shut down: %v", err)
	}

	if databases.DB != nil {
		err := databases.DB.Close()
		if err != nil {
			fmt.Printf("Error closing the database: %s\n", err.Error())
		} else {
			fmt.Println("Database connection closed successfully.")
		}
	}

	fmt.Println("Application stopped cleanly.")
}
