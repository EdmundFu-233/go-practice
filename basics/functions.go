package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(add(3, 4))
	a, b := swap(1, 2)
	fmt.Println(a, b)
}
