package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", ":8080")
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, client, _ := conn.ReadFromUDP(buf)
		fmt.Printf("Received %s from %s\n", string(buf[:n]), client)
		conn.WriteToUDP([]byte("OK\n"), client)
	}
}
