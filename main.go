package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/books/v1"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	// Manage your books
	BooksScope = "https://www.googleapis.com/auth/books"
)

func connectDB() {
	fmt.Println("Go now running")

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/go_googlebooks")
	if err != nil {
		fmt.Println("Something is not working")
	}
	defer db.Close()

	create, err := db.Query(`CREATE TABLE IF NOT EXISTS books (
		id INT(10) NOT NULL PRIMARY KEY AUTO_INCREMENT,
		book_title TEXT NOT NULL,
		book_desc TEXT NOT NULL,
		page_count INT (10) NOT NULL,
		publisher_id INT (10) NOT NULL
		)`)
	if err != nil {
		panic(err.Error())
	}

	create, err = db.Query(`CREATE TABLE IF NOT EXISTS authors (
		id INT(10) NOT NULL PRIMARY KEY AUTO_INCREMENT,
		author_name TEXT NOT NULL
	)`)

	if err != nil {
		panic(err.Error())
	}

	create, err = db.Query(`CREATE TABLE IF NOT EXISTS publishers (
		id INT(10) NOT NULL PRIMARY KEY AUTO_INCREMENT,
		book_title TEXT NOT NULL,
		author_name TEXT NOT NULL,
		publisher_id INT (10) NOT NULL
		)`)
	if err != nil {
		panic(err.Error())
	}

	defer create.Close()
}

// func getBooks(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("getting books")
// }

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	//gets all books in db
	myRouter.HandleFunc("/books", getBooks).Methods("GET")

	//get individual book
	myRouter.HandleFunc("/book{id}", getOneBook).Methods("GET")

	// Running the Server
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	ctx := context.Background()
	booksService, err := books.NewService(ctx)
	if err != nil {
		fmt.Println("something wrong with google books")
	}
	connectDB()
}
