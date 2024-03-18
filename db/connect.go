package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnect() *sql.DB {

	database, err := sql.Open("mysql", getDSN())
	if err != nil {
		panic(err.Error())
	}

	err = database.Ping()
	if err != nil {
		panic(err.Error())
	}

	return database
}

func getDSN() string {

	dbHost := os.Getenv("DB_HOST_KEY")
	dbPort := os.Getenv("DB_PORT_KEY")
	dbUser := os.Getenv("DB_USER_KEY")
	dbPassword := os.Getenv("DB_PASSWORD_KEY")
	dbName := os.Getenv("DB_NAME_KEY")
	path := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	fmt.Println(path)
	return path
}
