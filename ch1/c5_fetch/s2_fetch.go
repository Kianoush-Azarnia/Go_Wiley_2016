// Fetch prints the content found at a URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func fetch() {
	path := os.Args[0]
	fmt.Println("Path:", path)

	for _, url := range os.Args[1:] {
		http_prefix := "http://"
		// Excersice 1.8
		if !strings.HasPrefix(url, http_prefix) {
			url = http_prefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		body := new(strings.Builder)
		code1, err := io.Copy(body, resp.Body)

		fmt.Println(resp.Status, body.String()) // Excercise 1.9
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Sprintf("%d", code1)

	}
}
