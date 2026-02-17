package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func main() {
	http.HandleFunc("/", logging(hello))
	fmt.Println("Server on :8080")
	http.ListenAndServe(":8080", nil)
}
