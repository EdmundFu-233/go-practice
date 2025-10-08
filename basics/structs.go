package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Zhihao", Age: 22}
	fmt.Println(p.Name, p.Age)
}
