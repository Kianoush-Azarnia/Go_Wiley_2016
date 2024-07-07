package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func Dup1() {
	// streaming mode: read from stdin input into lines
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}
	// note: we ignored potential errors of input.Err()
	for line, num := range counts {
		if num > 1 {
			fmt.Printf("%d\t%s\n", num, line)
		}
	}
	// Q: when does this program stop reading from std input? EOF?
}

func Dup2() {
	// Dup2 reads from stdin or a list of named files
	// streaming mode: read from file or input into lines
	// streaming mode can handle any number of files or input lines
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// note: ignored potential errors of input.Err()
}

func Dup3() {
	// an alternative approach to streaming mode
	// read entire input into memory at once
	// split it into lines all at once then process
	counts := make(map[string]int)
	for _, fileName := range os.Args[1:] {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
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

// Exercise 1.4: modify Dup2 to print name of files with duplicate lines
func Dup2Files() {
	hasDup := make(map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Printf("No file is given.\n")
	} else {
		for _, arg := range files {
			counts := make(map[string]int)
			f, err := os.Open(arg)
			defer f.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			hasDup[arg] = HasDup(f, counts)
		}
	}
	for fileName, hasDupLine := range hasDup {
		if hasDupLine {
			fmt.Printf("%s\n", fileName)
		}
	}
}

func HasDup(f *os.File, counts map[string]int) bool {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if counts[line] > 1 {
			return true
		}
	}
	return false
	// note: ignored potential errors of input.Err()
}
