package main

import (
	"log"
	"os"

	file "linear-stats/filescanner"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage:go run . data.txt")
	}
	fileName := os.Args[1]
	file.InputScanner(fileName)
}
