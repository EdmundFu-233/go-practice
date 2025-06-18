package main

import (
	"fmt"
	"os"
	"strconv"
)

func celsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: converter <value> <unit>")
		return
	}
	value, _ := strconv.ParseFloat(os.Args[1], 64)
	unit := os.Args[2]
	switch unit {
	case "c2f":
		fmt.Printf("%.2f°C = %.2f°F\n", value, celsiusToFahrenheit(value))
	case "f2c":
		fmt.Printf("%.2f°F = %.2f°C\n", value, fahrenheitToCelsius(value))
	default:
		fmt.Println("Unknown unit")
	}
}
