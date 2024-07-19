package main

import "fmt"

// pc[i] is the population count of i
var pc [256]byte

func main() {
	// fmt.Println(pc)
	fmt.Println(PopCount(1021))
}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x
func PopCount(x uint64) int {
	var y int
	for i := 0; i < 8; i++ {
		// fmt.Println(x>>(i*8), byte(x>>(i*8)))
		y += int(pc[byte(x>>(i*8))])
	}
	return y
}
