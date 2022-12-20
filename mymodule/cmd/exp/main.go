package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	const (
		host     = "localhost"
		port     = 5432
		user     = "baloo"
		password = "junglebook"
		dbname   = "lenslocked"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// db, err := sql.Open("pgx", "host=localhost port=5432 user=hiteshkumar password=hitesh dbname=mymodule sslmode=disable")
	db, err := sql.Open("pgx", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected")
}
