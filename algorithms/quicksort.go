package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	var left, right []int
	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	result := append(quickSort(left), pivot)
	result = append(result, quickSort(right)...)
	return result
}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println(quickSort(arr))
}
