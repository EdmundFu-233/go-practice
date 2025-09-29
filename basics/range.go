package main

import "fmt"

func main() {
	nums := []int{2, 4, 6, 8}
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println("Sum:", sum)
	ages := map[string]int{"a": 1, "b": 2}
	for k, v := range ages {
		fmt.Printf("%s -> %d\n", k, v)
	}
}
