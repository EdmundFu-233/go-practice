package main

import "fmt"

func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	slice := []int{10, 20, 30}
	slice = append(slice, 40)
	fmt.Println(arr, slice)
}
