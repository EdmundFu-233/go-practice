package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

func handleProxy(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), 502)
		return
	}
	defer resp.Body.Close()
	for k, v := range resp.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	http.HandleFunc("/", handleProxy)
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Proxy on :8080")
	server.ListenAndServe()
}
