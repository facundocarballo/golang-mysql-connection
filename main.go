package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/facundocarballo/golang-mysql-connection/db"
	"github.com/facundocarballo/golang-mysql-connection/types"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const SERVER_PORT string = ":3690"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading the .env file.")
	}

	// Open connection to Database
	db, err := sql.Open("mysql", db.GetDSN())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Check success connection to Database
	if err = db.Ping(); err != nil {
		panic(err.Error())
	}

	// Define handlers to endpoints
	http.HandleFunc("/user", types.HandleUser)

	println("Server listening on port" + SERVER_PORT + " ...")
	http.ListenAndServe(SERVER_PORT, nil)
}
