package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// convertTemperature returns the result of converting from Celsius to Fahrenheit or from Fahrenheit to Celsius
// Usage rule: the inputUnit argument must have one of two possible values: 'F' or 'C'
func convertTemperature(inputValue float64, inputUnit string) (float64, string, error) {
	if inputUnit == "C" {
		return inputValue*9/5 + 32.0, "F", nil
	}

	if inputUnit == "F" {
		return (inputValue - 32) * 5 / 9, "C", nil
	}

	return 0, "", fmt.Errorf("%q: invalid input unit", inputUnit)
}

func main() {
	args := os.Args[1:]
	const errorMsg = "usage: tempconv <value> <C|F>"

	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "%s\n", errorMsg)
		os.Exit(1)
	}

	inputValue, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%q: %s\n", args[0], "invalid first input")
		os.Exit(1)
	}

	inputUnit := strings.ToUpper(args[1])

	convValue, convUnit, err := convertTemperature(inputValue, inputUnit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%.1f°%s = %.1f°%s\n", inputValue, inputUnit, convValue, convUnit)
}
