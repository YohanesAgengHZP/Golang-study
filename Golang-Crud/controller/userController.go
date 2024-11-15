package controller

import (
	"golang-crud/models"
	"text/template"
	"net/http"
)

//always implement error handling using panic(err) or customize handler

func Index(w http.ResponseWriter, r *http.Request) {
	users := models.GetAll()
	
	// initiate data variabel with keytype map and blank interface or any
	data := map[string]any{
		"users": users,
	}
	temp, err := template.ParseFiles("views/user/index.html")
	
	if err != nil {
		panic(err)	
	}

	temp.Execute(w, data)
}

// func Edit(w http.ResponseWriter, r *http.Request) {

// }

// func Delete(w http.ResponseWriter, r *http.Request) {

// }
