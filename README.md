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