package main

import (
	"fmt"
	"golang-crud/config"
	controllers "golang-crud/controller"
	"golang-crud/initializer"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVar()
	config.ConnectDatabase()
}

func main() {
	port := os.Getenv("PORT")
	r := gin.Default()
	r.POST("/add", controllers.CreateUser)
	r.GET("/users", controllers.ShowUser)
	r.GET("/users/:id", controllers.ShowUserByID)
	r.PUT("/users/:id", controllers.UpdateUserByID)
	r.DELETE("/users/:id", controllers.DeleteUserByID)

	r.Run()
	fmt.Println("Server running on port: ", port)
}
