package types

import "net/http"

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
func GetAllUsers() []User {

}

func GetUser(id int) User {

}

func CreateUser(name string, email string, password string) User {

}

func HandleUser(w http.ResponseWriter, r *http.Request) {
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
