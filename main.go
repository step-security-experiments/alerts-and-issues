package main

import "os"

func main() {
	_, _ = os.Open("example.txt") // This should ideally be flagged
}
