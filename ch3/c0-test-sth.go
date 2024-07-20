package main

import (
	"fmt"
)

func main() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Printf("%[2]d %[1]x %#[2]x %#[1]X\n", x, x+1)

	ascii := 'a'
	unicode := '漢'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 D 'D'"
	fmt.Printf("%d %[1]c %[1]q\n", newline)
}
