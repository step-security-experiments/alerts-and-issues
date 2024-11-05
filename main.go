package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Starting")

	// Example of a SQL injection vulnerability
	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		os.Exit(1)
	}
	defer db.Close()

	userInput := "1 OR 1=1" // Simulated unsafe user input
	query := fmt.Sprintf("SELECT * FROM users WHERE id = %s", userInput)
	_, err = db.Query(query) // This should trigger a CodeQL alert
	if err != nil {
		os.Exit(1)
	}

	os.Exit(1)
}
