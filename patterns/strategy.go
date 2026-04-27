package main

import "fmt"

type SortStrategy interface {
	Sort([]int) []int
}

type BubbleSort struct{}

func (b BubbleSort) Sort(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	copy(result, arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

type QuickSortS struct{}

func (q QuickSortS) Sort(arr []int) []int {
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
	result := append(q.Sort(left), pivot)
	result = append(result, q.Sort(right)...)
	return result
}

type Context struct {
	strategy SortStrategy
}

func (c *Context) Execute(arr []int) []int {
	return c.strategy.Sort(arr)
}

func main() {
	arr := []int{3, 1, 4, 1, 5}
	ctx := Context{strategy: BubbleSort{}}
	fmt.Println(ctx.Execute(arr))
}
