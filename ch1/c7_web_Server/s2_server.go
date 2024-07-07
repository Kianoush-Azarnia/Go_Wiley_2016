package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// server1 is a minimal echo server
func server1() {
	http.HandleFunc("/", handler1) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler1 echoes the path component of the requested url
func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// server2 is a minimal echo and counter server
var mu sync.Mutex
var count int

func server2() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter2)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler2 echoes the path component of the requested URL
func handler2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter2 echoes the number pf calls so far
func counter2(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// server3 calls handler3
func server3() {
	http.HandleFunc("/", handler3) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler3 echoes the HTTP request.
func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
