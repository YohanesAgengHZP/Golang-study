package utils

import (
	"net/http"
)

func HeaderHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
}