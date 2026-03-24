package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func scanPort(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
	if err != nil {
		return
	}
	conn.Close()
	fmt.Printf("Port %d open\n", port)
}

func main() {
	var wg sync.WaitGroup
	host := "localhost"
	for port := 1; port <= 1024; port++ {
		wg.Add(1)
		go scanPort(host, port, &wg)
	}
	wg.Wait()
}
