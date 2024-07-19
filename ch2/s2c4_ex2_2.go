package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func feetToMeters(feet float64) float64 {
	return feet * 0.3048
}

func metersToFeet(meters float64) float64 {
	return meters / 0.3048
}

func poundsToKilograms(pounds float64) float64 {
	return pounds * 0.453592
}

func kilogramsToPounds(kilograms float64) float64 {
	return kilograms / 0.453592
}

func convertTemperature(value float64, unit string) (float64, error) {
	unit = strings.ToLower(unit)
	switch unit {
	case "c":
		return celsiusToFahrenheit(value), nil
	case "f":
		return fahrenheitToCelsius(value), nil
	default:
		return 0, fmt.Errorf("Invalid unit. Supported units: C, F")
	}
}

func convertLength(value float64, unit string) (float64, error) {
	unit = strings.ToLower(unit)
	switch unit {
	case "ft":
		return feetToMeters(value), nil
	case "m":
		return metersToFeet(value), nil
	default:
		return 0, fmt.Errorf("Invalid unit. Supported units: ft, m")
	}
}

func convertWeight(value float64, unit string) (float64, error) {
	unit = strings.ToLower(unit)
	switch unit {
	case "lb":
		return poundsToKilograms(value), nil
	case "kg":
		return kilogramsToPounds(value), nil
	default:
		return 0, fmt.Errorf("Invalid unit. Supported units: lb, kg")
	}
}

func convertValue(value float64, unit string, convertType string) (float64, error) {
	convertType = strings.ToLower(convertType)
	switch convertType {
	case "temperature":
		return convertTemperature(value, unit)
	case "length":
		return convertLength(value, unit)
	case "weight":
		return convertWeight(value, unit)
	default:
		return 0, fmt.Errorf("Invalid conversion type. Supported types: temperature, length, weight")
	}
}

func main() {
	if len(os.Args) > 1 {
		// Conversion from command-line arguments
		for i := 1; i < len(os.Args); i += 3 {
			value, err := strconv.ParseFloat(os.Args[i], 64)
			if err != nil {
				fmt.Println("Invalid input format. Please provide value, unit, and conversion type.")
				return
			}
			unit := os.Args[i+1]
			convertType := os.Args[i+2]
			convertedValue, err := convertValue(value, unit, convertType)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%.2f %s is %.2f %s\n", value, unit, convertedValue, convertType)
		}
	} else {
		// Conversion from standard input
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("Enter value, unit, and conversion type (separated by space): ")
			if !scanner.Scan() {
				break
			}
			inputStr := scanner.Text()
			if inputStr == "" {
				break
			}
			inputs := strings.Split(inputStr, " ")
			if len(inputs) != 3 {
				fmt.Println("Invalid input format. Please provide value, unit, and conversion type.")
				continue
			}
			value, err := strconv.ParseFloat(inputs[0], 64)
			if err != nil {
				fmt.Println("Invalid input format. Please provide value, unit, and conversion type.")
				continue
			}
			unit := inputs[1]
			convertType := inputs[2]
			convertedValue, err := convertValue(value, unit, convertType)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("%.2f %s is %.2f %s\n", value, unit, convertedValue, convertType)
		}
	}
}
