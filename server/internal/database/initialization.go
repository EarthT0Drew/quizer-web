package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func OpenConnection(dbName string) *sql.DB {
	var host string = "database"
	var user string = os.Getenv("POSTGRES_USER")
	var password string = os.Getenv("POSTGRES_PASSWORD")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%v user=%v password=%v port=5432 sslmode=disable dbname=%v", host, user, password, dbName))

	err = db.Ping() // Ping will return an error if we can't authenticate

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully authenticated with the database '" + dbName + "'.")
	}

	return db
}
