package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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

// just an example of reading a file
func Dup2() {
	// Dup2 reads from stdin or from a list of named files
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
