package main

import "fmt"

func main() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a' // makes uppercase?
			// book: above ie not equivalent to unicode.ToUpper
			fmt.Printf("%c", x) // HELLO (one char per iterate)
		}
	}
}
