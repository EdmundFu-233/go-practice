package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		for {
			_ = make([]byte, 1024*1024)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	fmt.Println("Profiling on :6060")
	http.ListenAndServe(":6060", nil)
}
