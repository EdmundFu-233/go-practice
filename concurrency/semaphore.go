// Package concurrency provides concurrency pattern implementations.
package concurrency

import (
	"context"
	"sync"
	"time"
)

// Semaphore implements a weighted semaphore using a buffered channel.
// It limits concurrent access to a resource to a fixed number of goroutines.
type Semaphore struct {
	tickets chan struct{}
}

// NewSemaphore creates a semaphore with the given maximum concurrency.
func NewSemaphore(maxConcurrency int) *Semaphore {
	if maxConcurrency <= 0 {
		maxConcurrency = 1
	}
	return &Semaphore{
		tickets: make(chan struct{}, maxConcurrency),
	}
}

// Acquire blocks until a permit is available or ctx is cancelled.
func (s *Semaphore) Acquire(ctx context.Context) error {
	select {
	case s.tickets <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// TryAcquire attempts to acquire a permit without blocking.
// Returns true if the permit was acquired, false otherwise.
func (s *Semaphore) TryAcquire() bool {
	select {
	case s.tickets <- struct{}{}:
		return true
	default:
		return false
	}
}

// Release returns a permit to the semaphore.
// Panics if Release is called more times than Acquire.
func (s *Semaphore) Release() {
	select {
	case <-s.tickets:
	default:
		panic("semaphore: Release called more times than Acquire")
	}
}

// Available returns the number of available permits.
func (s *Semaphore) Available() int {
	return cap(s.tickets) - len(s.tickets)
}

// Capacity returns the maximum concurrency of the semaphore.
func (s *Semaphore) Capacity() int {
	return cap(s.tickets)
}

// Wait blocks until all acquired permits are released (drains the channel).
func (s *Semaphore) Wait() {
	for len(s.tickets) > 0 {
		<-s.tickets
	}
}

// ---- Weighted Semaphore ----

// WeightedSemaphore is a semaphore that supports acquiring weighted permits.
// Each Acquire requests a specific number of permits, and will block until
// enough permits are available.
type WeightedSemaphore struct {
	mu      sync.Mutex
	cond    *sync.Cond
	current int64
	max     int64
}

// NewWeightedSemaphore creates a weighted semaphore with total capacity.
func NewWeightedSemaphore(max int64) *WeightedSemaphore {
	if max <= 0 {
		max = 1
	}
	s := &WeightedSemaphore{max: max}
	s.cond = sync.NewCond(&s.mu)
	return s
}

// Acquire blocks until n permits are available.
func (s *WeightedSemaphore) Acquire(ctx context.Context, n int64) error {
	if n <= 0 {
		return nil
	}
	if n > s.max {
		return ErrAcquireExceedsMax
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for {
		if s.current+n <= s.max {
			s.current += n
			return nil
		}
		// Wait with context cancellation
		waitChan := make(chan struct{})
		go func() {
			s.cond.Wait()
			close(waitChan)
		}()

		select {
		case <-waitChan:
			continue
		case <-ctx.Done():
			// Signal to wake the goroutine
			s.cond.Signal()
			return ctx.Err()
		}
	}
}

// Release returns n permits to the semaphore.
func (s *WeightedSemaphore) Release(n int64) {
	if n <= 0 {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	s.current -= n
	if s.current < 0 {
		s.current = 0
	}
	s.cond.Broadcast()
}

// Available returns the current number of available permits.
func (s *WeightedSemaphore) Available() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.max - s.current
}

// ErrAcquireExceedsMax is returned when trying to acquire more permits than
// the semaphore's maximum capacity.
var ErrAcquireExceedsMax = &SemaphoreError{"acquire exceeds maximum capacity"}

// SemaphoreError represents a semaphore-specific error.
type SemaphoreError struct{ msg string }

func (e *SemaphoreError) Error() string { return e.msg }

// ---- Rate Limiter using Semaphore ----

// RateLimiter limits the rate of operations using a token bucket approach.
type RateLimiter struct {
	sem     *Semaphore
	ticker  *time.Ticker
	stopCh  chan struct{}
	started bool
	mu      sync.Mutex
}

// NewRateLimiter creates a rate limiter allowing rate operations per second.
func NewRateLimiter(rate int) *RateLimiter {
	if rate <= 0 {
		rate = 1
	}
	return &RateLimiter{
		sem:    NewSemaphore(rate),
		stopCh: make(chan struct{}),
	}
}

// Start begins periodically adding permits.
func (rl *RateLimiter) Start() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if rl.started {
		return
	}
	rl.started = true

	// Refill one permit per interval
	interval := time.Second / time.Duration(cap(rl.sem.tickets))
	rl.ticker = time.NewTicker(interval)

	go func() {
		for {
			select {
			case <-rl.ticker.C:
				select {
				case <-rl.sem.tickets:
					// Drain one ticket to keep it available
				default:
				}
			case <-rl.stopCh:
				rl.ticker.Stop()
				return
			}
		}
	}()
}

// Stop stops the rate limiter refill.
func (rl *RateLimiter) Stop() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if !rl.started {
		return
	}
	rl.started = false
	close(rl.stopCh)
}

// Allow checks if an operation is allowed right now.
func (rl *RateLimiter) Allow() bool {
	return rl.sem.TryAcquire()
}

// Wait blocks until a permit is available.
func (rl *RateLimiter) Wait(ctx context.Context) error {
	return rl.sem.Acquire(ctx)
}
