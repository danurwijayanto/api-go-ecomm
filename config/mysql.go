package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// GetMysqlConnect : function for connect to mysql databases
func GetMysqlConnect() (*sql.DB, error) {
	dbUser := "root"
	dbPass := ""
	dbName := "pulsa_ecomm"
	dbHost := "localhost"
	dbPort := "3306"

	uri := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

	db, err := sql.Open("mysql", uri)

	if err != nil {
		log.Fatal(err)
	}

	return db, err

}
