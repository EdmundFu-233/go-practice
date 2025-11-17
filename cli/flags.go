package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "World", "a name to greet")
	count := flag.Int("count", 1, "number of times to greet")
	flag.Parse()
	for i := 0; i < *count; i++ {
		fmt.Printf("Hello, %s!\n", *name)
	}
}
