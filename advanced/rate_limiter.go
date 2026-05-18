package advanced

import (
	"fmt"
	"sync"
	"time"
)

// TokenBucket implements a rate limiter using the token bucket algorithm.
type TokenBucket struct {
	rate       float64
	capacity   float64
	tokens     float64
	lastRefill time.Time
	mu         sync.Mutex
}

// NewTokenBucket creates a new token bucket with given rate (tokens/sec) and capacity.
func NewTokenBucket(rate, capacity float64) *TokenBucket {
	return &TokenBucket{
		rate:       rate,
		capacity:   capacity,
		tokens:     capacity,
		lastRefill: time.Now(),
	}
}

// Allow checks if a request is allowed, consuming one token if so.
func (tb *TokenBucket) Allow() bool {
	return tb.AllowN(1)
}

// AllowN checks if n requests are allowed, consuming n tokens if so.
func (tb *TokenBucket) AllowN(n float64) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(tb.lastRefill).Seconds()
	tb.tokens += elapsed * tb.rate
	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastRefill = now

	if tb.tokens >= n {
		tb.tokens -= n
		return true
	}
	return false
}

// LeakyBucket implements a rate limiter using the leaky bucket algorithm.
type LeakyBucket struct {
	capacity int64
	queue    chan struct{}
	rate     time.Duration
}

// NewLeakyBucket creates a leaky bucket that processes at most `rate` requests per second.
func NewLeakyBucket(capacity int64, ratePerSec int) *LeakyBucket {
	lb := &LeakyBucket{
		capacity: capacity,
		queue:    make(chan struct{}, capacity),
		rate:     time.Second / time.Duration(ratePerSec),
	}
	go lb.drip()
	return lb
}

func (lb *LeakyBucket) drip() {
	ticker := time.NewTicker(lb.rate)
	defer ticker.Stop()
	for range ticker.C {
		select {
		case <-lb.queue:
		default:
		}
	}
}

// Allow tries to add a request; returns false if bucket is full.
func (lb *LeakyBucket) Allow() bool {
	select {
	case lb.queue <- struct{}{}:
		return true
	default:
		return false
	}
}

func ExampleRateLimiter() {
	tb := NewTokenBucket(10, 5) // 10 tokens/sec, burst of 5
	for i := 0; i < 10; i++ {
		if tb.Allow() {
			fmt.Printf("TokenBucket: request %d allowed\n", i+1)
		} else {
			fmt.Printf("TokenBucket: request %d rate limited\n", i+1)
		}
		time.Sleep(50 * time.Millisecond)
	}
}
