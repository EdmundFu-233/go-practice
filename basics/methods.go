package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(f float64) {
	r.Width *= f
	r.Height *= f
}

func main() {
	r := Rectangle{3, 4}
	fmt.Println("Area:", r.Area())
	r.Scale(2)
	fmt.Println("After scale:", r.Area())
}
