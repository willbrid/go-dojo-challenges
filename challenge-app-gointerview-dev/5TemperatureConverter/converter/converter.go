package converter

import "math"

func CelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32.0
}

func FahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func Round(value float64, decimals int) float64 {
	precision := math.Pow10(decimals)

	return math.Round(value*precision) / precision
}
