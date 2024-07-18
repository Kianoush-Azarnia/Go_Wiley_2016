package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}
}

func fibonacci(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
