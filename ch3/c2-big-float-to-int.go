package main

import (
	"fmt"
)

func main() {
	f := 1e100
	fmt.Printf("f: %T, %v\n", f, f)

	i := int(f)
	fmt.Printf("i: %T, %v\n", i, i)
}
