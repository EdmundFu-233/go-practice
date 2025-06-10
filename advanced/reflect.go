package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	Field1 string
	Field2 int
}

func main() {
	s := MyStruct{"hello", 42}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	fmt.Println("Type:", t.Name())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("  %s = %v\n", field.Name, value.Interface())
	}
}
