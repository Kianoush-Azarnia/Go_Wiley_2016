package main

import "fmt"

func main() {
	x := "hello!"
	for _, x := range x {
		if x != '!' {
			x := x + 'A' - 'a' // makes uppercase?
			// book: above ie not equivalent to unicode.ToUpper
			fmt.Printf("%c", x) // HELLO (one char per iterate)
		}
	}
}
