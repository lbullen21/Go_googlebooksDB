package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go now running")

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/go_googlebooks")
	if err != nil {
		fmt.Println("Something is not working")
	}
	defer db.Close()

	create, err := db.Query("CREATE TABLE Publishers (Id INT PRIMARY KEY, Publishers VARCHAR(255), Title VARCHAR(255), Author VARCHAR(255))")

	if err != nil {
		panic(err.Error())
	}

	defer create.Close()
}
