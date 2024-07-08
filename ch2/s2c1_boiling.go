// Boiling prints the boiling pint of water
package main

import "fmt"

const boilingF = 212.0

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g F = %g C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g F = %g C\n", boilingF, fToC(boilingF))

	var p = fu()
	fmt.Printf("%v\n", *p)
	fmt.Printf("%v\n", fu() == fu())
}

func printBoiling() {
	var f = boilingF
	var c = (f - 32) * 5 / 9

	fmt.Printf("boiling point = %g F or %g C\n", f, c)
	// Output: boiling point = 212 F or 100 C
}

// Ftoc prints two Farenheit-to-Celsius conversions.
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

// playing with pointers
func fu() *int {
	v := 1
	return &v
}
