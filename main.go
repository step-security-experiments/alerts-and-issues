package main

import (
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("example.txt") // This should raise an alert: ignoring errors
	defer file.Close()
	fmt.Println("File opened")
}
