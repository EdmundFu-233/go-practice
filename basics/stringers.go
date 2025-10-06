package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years)", p.Name, p.Age)
}

func main() {
	p := Person{"Zhihao", 22}
	fmt.Println(p)
}
