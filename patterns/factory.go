package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}

func (c Cat) Speak() string { return "Meow!" }

func AnimalFactory(animalType string) Animal {
	switch animalType {
	case "dog":
		return Dog{}
	case "cat":
		return Cat{}
	default:
		return nil
	}
}

func main() {
	dog := AnimalFactory("dog")
	cat := AnimalFactory("cat")
	fmt.Println(dog.Speak())
	fmt.Println(cat.Speak())
}
