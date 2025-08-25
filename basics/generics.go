package main

import "fmt"

func Map[T any](s []T, f func(T) T) []T {
	result := make([]T, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	doubled := Map(nums, func(n int) int { return n * 2 })
	fmt.Println(doubled)

	words := []string{"a", "b", "c"}
	upper := Map(words, func(s string) string { return s + "!" })
	fmt.Println(upper)
}
