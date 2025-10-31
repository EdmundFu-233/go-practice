package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No arguments provided")
		return
	}
	for i, arg := range args {
		fmt.Printf("arg[%d] = %s\n", i, arg)
	}
}
