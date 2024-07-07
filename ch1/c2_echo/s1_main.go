// run code commands:
// first time: just run first line of this code
// after that: xargs -d '\n' go run s1_main.go s2_echo.go < echo_args.txt
package main

import "fmt"

func main() {
	// WriteInFile("echo_args.txt", "directory", 10000)

	functions := []func(){
		Echo_1,
		Echo_2,
		Echo_3,
		Echo_4,
		Echo_5,
		Echo_6,
	}

	fmt.Println(Run_And_Measure(functions))
}
