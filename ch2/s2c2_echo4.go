package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	Echo4()
}

func Echo4() {
	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", " ", "separator")

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}
