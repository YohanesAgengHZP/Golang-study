package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	databases "generic_crud/database"
	"generic_crud/entities"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// GetUsers fetches all user data
func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all user data")

	var users []entities.User

	rows, err := databases.DB.Query("SELECT id, name FROM users")
	if err != nil {
		http.Error(w, "Failed to fetch users: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			http.Error(w, "Failed to scan user: "+err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByID fetches a user by ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Fetching user by ID")
    vars := mux.Vars(r)
    idParam := vars["id"]

    if idParam == "" {
        http.Error(w, "ID parameter is missing", http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(idParam)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    query := "SELECT id, name, email FROM users WHERE id = ?"
    var user entities.User
    err = databases.DB.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, "Failed to fetch user: "+err.Error(), http.StatusInternalServerError)
        }
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}


// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating a new user")

	// Decode JSON Request Body
	var user entities.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Input validation
	if user.Name == "" || user.Password == "" || user.Email == "" {
		http.Error(w, "Missing required fields: name, password, or email", http.StatusBadRequest)
		return
	}

	// Insert uery
	query := "INSERT INTO users (name, password, email) VALUES (?, ?, ?)"
	result, err := databases.DB.Exec(query, user.Name, user.Password, user.Email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Retrieve the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve user ID", http.StatusInternalServerError)
		return
	}

	// Set the generated ID and local timestamp
	user.Id = int(id)
	user.CreatedAt = time.Now().Local()

	// Respond with the created user data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}


// UpdateUser updates a user's data
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Updating user")

    vars := mux.Vars(r)
    idParam := vars["id"]
    id, err := strconv.Atoi(idParam)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    var user entities.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
        return
    }

    query := "UPDATE users SET name = ?, email = ? WHERE id = ?"
    _, err = databases.DB.Exec(query, user.Name, user.Email, id)
    if err != nil {
        http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
        return
    }

    user.Id = id

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Deleting user")

    vars := mux.Vars(r)
    idParam := vars["id"]
    id, err := strconv.Atoi(idParam)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    query := "DELETE FROM users WHERE id = ?"
    _, err = databases.DB.Exec(query, id)
    if err != nil {
        http.Error(w, "Failed to delete user: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
