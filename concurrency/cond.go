package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	ready := false

	go func() {
		time.Sleep(100 * time.Millisecond)
		mu.Lock()
		ready = true
		cond.Broadcast()
		mu.Unlock()
	}()

	mu.Lock()
	for !ready {
		cond.Wait()
	}
	fmt.Println("Ready!")
	mu.Unlock()
}
