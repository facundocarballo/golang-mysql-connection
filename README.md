# Golang - MySQL - Connection

## Step by Step - Server Setup
1. Init Golang Project
```bash
go mod init github.com/facundocarballo/golang-mysql-connection
```
2. Init Github Repo
```bash
git init
```
3. Install GoDotEnv for .env files
```bash
    go get github.com/joho/godotenv
```
4. Install Go-SQL-Driver
```bash
    go get github.com/go-sql-driver/mysql
```
> It's important to know that you will have to import this library to connect your golang server with your mysql database. Normally using like this example.
```golang
    import (
        _ "github.com/go-sql-driver/mysql"
    )
```
5. Run the server
```bash
    go run main.go
```

---

## API Calls - Examples
### User
This endpoint works for create a new user, and also for getting all the users in our database.
#### POST REQUEST
To create a new user, you have to call to this endpoint **'/user'** with a POST REQUEST and provide this data in the body of the request.
```json
    {
        "name": "Name of the user",
        "email": "email@user.com",
        "password": "strongPassword",
    }
```
Using curl will be something like this.
```bash
    curl -X POST \
    -H "Content-Type: application/json" \
    -d '{
        "name": "Name of the user",
        "email": "email@user.com",
        "password": "strongPassword"
    }' \
    http://localhost:3690/user
```
#### GET REQUEST
To get all the users that are stored in this database, you will have to call to this endpoint **'/user'** with a GET REQUEST.
Using curl will be something like this.
```bash
    curl http://localhost:3690/user
```