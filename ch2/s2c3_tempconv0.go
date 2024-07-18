// Package tempconv performs Celsius and Fahrenheit temperature computations
package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == Celsius(f))

	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F

	d := FToC(212.0)
	fmt.Println(d.String()) // "100°C"
	fmt.Printf("%v\n", d)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", d)   // "100°C"
	fmt.Println(d)          // "100°C"
	fmt.Printf("%g\n", d)   // "100"; does not call String
	fmt.Println(float64(d)) // "100"; does not call String
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
