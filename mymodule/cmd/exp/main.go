package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	db, err := sql.Open("pgx", "host=localhost port=5432 user=postgres password=postgres dbname=mymodule sslmode=disable")

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
