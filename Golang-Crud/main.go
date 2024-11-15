package main

import (
	"fmt"
	"golang-crud/config"
	homecontroller "golang-crud/controller"
	userController "golang-crud/controller"

	"net/http"
)

func main() {
	config.ConnectDB()

	//1 Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	//2. user
	http.HandleFunc("/users", userController.Index)
	// http.HandleFunc("/users/add", userController.App)
	// http.HandleFunc("/users/edit", userController.Edit)
	// http.HandleFunc("/users/delete", userController.Delete)

	port := "8080"
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server is listening on port:", port)
	}
}
