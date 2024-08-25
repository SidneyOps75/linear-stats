package reader

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	formula "linear-stats/formulas"
)

func InputScanner(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("failed to open file")
	}
	defer file.Close()

	var y []float64
	var x []float64

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		text := scanner.Text()
		number, err := strconv.ParseFloat(text, 64)
		if err != nil {
			log.Fatal("Invalid text to convert to float")
		}
		y = append(y, number)
		x = append(x, float64(lineNumber))
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file")
	}

	// Call LinearRegression with x as line numbers and y as the file data
	m, c := formula.LinearRegression(x, y)

	// Call PearsonCorrelation with x and y
	r := formula.PearsonCorrelation(x, y)

	// Print results with required precision
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, c)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
