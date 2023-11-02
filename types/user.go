package types

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) UpdateUser() bool {

}

func (u User) DeleteUser() bool {

}

// "Static" functions
func GetAllUsers(w http.ResponseWriter, database *sql.DB) []User {
	rows, err := database.Query("SELECT id, name, email FROM User")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Iterate Rows
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return nil
		}
		users = append(users, user)
	}

	// Check Error on Rows
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	// Send response to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetUser(id int) User {

}

func CreateUser(name string, email string, password string) User {

}

func HandleUser(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	if r.Method == "POST" {
		// Get the user data
		// Post the new user
		return
	}

	if r.Method == "GET" {
		// Check if pass parameters as id? -> GetUser : GetAllUsers
	}

	if r.Method == "UPDATE" {
		// Check params to update user.
	}

	if r.Method == "DELETE" {
		// Check params to delete user.
	}
}
