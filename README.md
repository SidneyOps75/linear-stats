# Linear Stats

This project calculates statistical values such as the Linear Regression Line and Pearson Correlation Coefficient from a dataset. The dataset consists of values for the y axis, with the x axis implicitly being the line numbers.
Features

Linear Regression Line Calculation: Calculates the slope (m) and intercept (b) of the line fitting the dataset.
Pearson Correlation Coefficient: Computes the correlation coefficient to show the strength and direction of the linear relationship between x and y.

## Project Structure

```bash

linear-stats/
├── data.txt                   # Sample dataset
├── main.go                    # Main program entry
├── go.mod                     # Go module definition
├── formulas/                  # Package for statistical formulas
│   ├── mean.go                # Mean calculation
│   ├── linear-regression.go   # Linear regression calculations
│   ├── pearson-correlation.go # Pearson correlation coefficient calculations
│   └── mean_test.go           # Unit tests for mean function
└── filescanner/
    └── file-scanner.go        # File scanner to read dataset
```

## Installation

1. Clone the repository:

```bash

git clone https://learn.zone01kisumu.ke/git/johopiyo/linear-stats
cd linear-stats
```

2. Install Go dependencies:

```bash

    go mod tidy
```
## Usage

1. To run the program, use the following command, passing the file containing the dataset:

```bash

go run main.go data.txt
```
2. Input File Format

The input file should contain one y value per line. The x values will be automatically generated as the line numbers.

Example data.txt file:

```bash
189
113
121
114
145
110
...
```

3. Output

The output will display the Linear Regression Line and Pearson Correlation Coefficient in the following format:

Linear Regression Line: y = <m>x + <b>
Pearson Correlation Coefficient: <r>

Where:

    <m> is the slope of the line.
    <b> is the y-intercept.
    <r> is the Pearson Correlation Coefficient.

## Testing

To run tests for the formulas, use:

```bash
go test ./formulas/...
```

## License

This project is licensed under the [MIT](./LICENSE) License.

## Author

This project has been authored by [John Opiyo](https://github.com/SidneyOps75)