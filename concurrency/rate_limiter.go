// Package concurrency implements reusable concurrency patterns in Go.
//
// Rate Limiter
//
// Implements a token-bucket rate limiter that controls the rate at which
// operations are allowed to proceed. Useful for API clients, throttling
// outgoing requests, or controlling resource consumption in concurrent
// workloads.
package concurrency

import (
	"sync"
	"time"
)

// RateLimiter implements a thread-safe token-bucket rate limiter.
//
// Tokens are added to the bucket at a fixed rate up to a maximum
// capacity. Each Allow() call consumes one token if available.
type RateLimiter struct {
	mu       sync.Mutex
	rate     float64    // tokens per second
	capacity float64    // max tokens in bucket
	tokens   float64    // current token count
	lastTime time.Time  // last token refill time
}

// NewRateLimiter creates a token-bucket rate limiter.
//
//   - rate: maximum sustained rate in operations per second
//   - capacity: burst size (max tokens accumulated when idle)
func NewRateLimiter(rate, capacity float64) *RateLimiter {
	return &RateLimiter{
		rate:     rate,
		capacity: capacity,
		tokens:   capacity, // start full
		lastTime: time.Now(),
	}
}

// Allow checks whether one operation is allowed under the rate limit.
// Returns true if a token was consumed, false if the rate limit has
// been exceeded.
func (rl *RateLimiter) Allow() bool {
	return rl.AllowN(1)
}

// AllowN checks whether n operations are allowed under the rate limit.
func (rl *RateLimiter) AllowN(n float64) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastTime).Seconds()

	// Refill tokens based on elapsed time
	rl.tokens += elapsed * rl.rate
	if rl.tokens > rl.capacity {
		rl.tokens = rl.capacity
	}
	rl.lastTime = now

	if rl.tokens >= n {
		rl.tokens -= n
		return true
	}
	return false
}

// Wait blocks until n operations are allowed, respecting the rate limit.
// Returns immediately if enough tokens are available; otherwise sleeps
// until the required tokens accumulate.
func (rl *RateLimiter) Wait(n float64) {
	for !rl.AllowN(n) {
		rl.mu.Lock()
		needed := n - rl.tokens
		rl.mu.Unlock()

		// Sleep for the time needed to accumulate missing tokens
		sleepDuration := time.Duration(needed / rl.rate * float64(time.Second))
		if sleepDuration < time.Millisecond {
			sleepDuration = time.Millisecond
		}
		time.Sleep(sleepDuration)
	}
}

// Tokens returns the current number of available tokens (approximate).
func (rl *RateLimiter) Tokens() float64 {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastTime).Seconds()
	tokens := rl.tokens + elapsed*rl.rate
	if tokens > rl.capacity {
		tokens = rl.capacity
	}
	return tokens
}

// SetRate updates the refill rate (tokens/second).
func (rl *RateLimiter) SetRate(rate float64) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	rl.rate = rate
}
