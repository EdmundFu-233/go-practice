package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: dns_lookup <hostname>")
		return
	}
	hostname := os.Args[1]
	ips, err := net.LookupHost(hostname)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
