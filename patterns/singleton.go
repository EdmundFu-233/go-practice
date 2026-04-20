package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	value string
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{value: "initialized"}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()
	fmt.Println(s1 == s2)
}
