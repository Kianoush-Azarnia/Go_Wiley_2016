// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := Fahrenheit(t)
		c := Celsius(t)

		fmt.Printf("%v = %v, %v = %v\n", f, FToC(f), c, CToF(c))
		fmt.Printf("%v = %v, %v = %v\n", f, FToK(f), c, CToK(c))
	}
}

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f-32)*5/9 + 273.15)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%.2f°K", k)
}
