package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open an in-memory SQLite database
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the database connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//set id = environment variable exampleid
	id := os.Getenv("EXAMPLE_ID")

	// Create a new table
	sqlStmt := `SELECT * FROM DATE_TABLE WHERE id = ` + id

	//execute the query
	_, err = db.Query(sqlStmt)

	currentDate := time.Now().Format("2006-01-02")
	fmt.Println("Current date:", currentDate)
}
