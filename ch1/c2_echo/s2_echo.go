package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func Echo_1() {
	var s string         // implicit initialization
	var sep string = " " // explicit init
	for i := 1; i < len(os.Args); i++ {
		s = os.Args[i]
		_ = fmt.Sprintf("%s%s", s, sep)
	}
	fmt.Println()
}

func Echo_2() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		// s += os.Args[i] + sep  wrong (but why?)
		// because code initiates the sep inside the loop (next line)
		sep = " "
	}
	_ = fmt.Sprintf("%s\n", s)
}

func Echo_3() {
	s, sep := "", "" // two short variable declarations together
	// here for iterates over a range: range returns index, value
	for _, arg := range os.Args[1:] {
		// _ underscore is blank identifier (unused) for index
		s += sep + arg
		sep = " "
	}
	_ = fmt.Sprintf("%s\n", s)
}

func Echo_4() {
	// more efficient for large data, because it doesn't re-assign in loop
	s := strings.Join(os.Args[1:], " ")
	_ = fmt.Sprintf("%s\n", s)
}

func Echo_5() {
	// Excercise 1.1: print command as well
	_ = fmt.Sprintf("%s", os.Args[1:])
	_ = fmt.Sprintf("command: %s\n", os.Args[0]) // command
}

func Echo_6() {
	sep := ", "
	args := os.Args[1:]
	// Excercise 1.2: print index vlaue
	for idx, val := range args {
		if idx == len(args)-1 {
			fmt.Printf("%v: %s", idx, val)
			break
		}
		_ = fmt.Sprintf("%v: %s%s", idx, val, sep)
	}
	fmt.Println("")
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

func WriteInFile(fileName, expression string, linesNum int) {
	// Open a file for writing
	file, err := os.Create(fmt.Sprintf("%s", fileName))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write lines to the file
	for i := 1; i <= linesNum; i++ {
		_, err := file.WriteString(fmt.Sprintf("%s%d\n", expression, i))
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("File written successfully")
}
