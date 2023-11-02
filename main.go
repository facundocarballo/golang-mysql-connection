package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const DB_HOST_KEY = "DB_HOST"
const DB_PORT_KEY = "DB_PORT"
const DB_PASSWORD_KEY = "DB_PASSWORD"
const DB_USER_KEY = "DB_USER"
const DB_NAME_KEY = "DB_NAME"

const SERVER_PORT string = ":3690"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading the .env file.")
	}

	// Get Datbase data from .env file
	db_host := os.Getenv(DB_HOST_KEY)
	db_port := os.Getenv(DB_PORT_KEY)
	db_user := os.Getenv(DB_USER_KEY)
	db_password := os.Getenv(DB_PASSWORD_KEY)
	db_name := os.Getenv(DB_NAME_KEY)

	dsn := db_user + ":" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + db_name

	// Open connection to Database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Check success connection to Database
	if err = db.Ping(); err != nil {
		panic(err.Error())
	}

	// Define handlers to endpoints
	http.HandleFunc("/user", HandlerUser)

	println("Server listening on port" + SERVER_PORT + " ...")
	http.ListenAndServe(SERVER_PORT, nil)
}

func HandlerUser(w http.ResponseWriter, r *http.Request) {

}
