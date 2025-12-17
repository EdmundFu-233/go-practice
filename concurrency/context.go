package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Work done")
	case <-ctx.Done():
		fmt.Println("Cancelled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	go doWork(ctx)
	time.Sleep(1 * time.Second)
}
