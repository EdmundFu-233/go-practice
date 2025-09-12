package main

import "fmt"

func main() {
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	ages["Charlie"] = 35
	delete(ages, "Bob")
	fmt.Println(ages)
}
