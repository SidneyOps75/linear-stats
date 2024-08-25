package reader

import (
	"bufio"
	"fmt"
	formula "linear-stats/formulas"
	"log"
	"os"
	"strconv"
)

func InputScanner(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("failed to open file")
	}
	defer file.Close()

	var y []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		number, err := strconv.ParseFloat(text, 64)
		if err != nil {
			log.Fatal("Invalid text to convert to float")
		}
		y = append(y, number)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file")
	}
	m, c := formula.LinearRegression(y)
	r := formula.PearsonCorrelation(y)
	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", m, c)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)

}
