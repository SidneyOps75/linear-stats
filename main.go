package main

import (
	"go/scanner"
	"os"
)

func main() {
	filename := os.Args[1]
	scanner.InputScanner(filename)
}
