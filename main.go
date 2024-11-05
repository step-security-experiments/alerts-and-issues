package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting")
	os.Exit(1) // This line will cause the CodeQL alert
}
