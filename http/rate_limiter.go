package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	mu      sync.Mutex
	visits  map[string]int
	limit   int
	window  time.Duration
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{visits: make(map[string]int), limit: limit, window: window}
}

func (rl *RateLimiter) Limit(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		rl.mu.Lock()
		rl.visits[ip]++
		count := rl.visits[ip]
		rl.mu.Unlock()
		if count > rl.limit {
			http.Error(w, "Rate limit exceeded", 429)
			return
		}
		next(w, r)
	}
}

func main() {
	rl := NewRateLimiter(10, time.Minute)
	http.HandleFunc("/", rl.Limit(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	}))
	fmt.Println("Server on :8080")
	http.ListenAndServe(":8080", nil)
}
