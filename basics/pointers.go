package main

import "fmt"

func zero(val *int) {
	*val = 0
}

func main() {
	i := 42
	p := &i
	fmt.Println("Before:", i)
	zero(p)
	fmt.Println("After:", i)
}
