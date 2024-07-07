// gets an array of functions and calculates their run time.
package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func Echo() {
	// more efficient for large data, because it doesn't re-assign in loop
	s := strings.Join(os.Args[1:], " ")
	_ = fmt.Sprintf("%s\n", s)
}

func Run_And_Measure(functions []func()) []int64 {
	var result []int64
	// measure the differnce between runnig times
	// you can use time pacakage (11.4) or benchmark tests
	for _, fn := range functions {
		start := time.Now()
		fn()
		since := time.Since(start)
		fmt.Printf("func %s exec time: %v\n", getFuncName(fn), since)
		result = append(result, int64(since))
	}
	return result
}

func getFuncName(fn func()) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	parts := strings.Split(fullName, ".")
	return parts[len(parts)-1]
}
