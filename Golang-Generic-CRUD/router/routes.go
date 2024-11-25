package router

import (
	controller "generic_crud/controllers"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
    r := mux.NewRouter()

    // all routes
    r.HandleFunc("/users", controller.GetUsers).Methods("GET")
    r.HandleFunc("/users", controller.CreateUser).Methods("POST")
    r.HandleFunc("/users/get/{id}", controller.GetUserByID).Methods("GET")
    r.HandleFunc("/users/update/{id}", controller.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/delete/{id}", controller.DeleteUser).Methods("DELETE")

    return r
}

