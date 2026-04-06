package main

import "fmt"

type Pizza struct {
	Size     string
	Cheese   bool
	Pepperoni bool
	Veggies  []string
}

type PizzaBuilder struct {
	size      string
	cheese    bool
	pepperoni bool
	veggies   []string
}

func NewPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{}
}

func (b *PizzaBuilder) SetSize(size string) *PizzaBuilder {
	b.size = size; return b
}

func (b *PizzaBuilder) AddCheese() *PizzaBuilder {
	b.cheese = true; return b
}

func (b *PizzaBuilder) AddPepperoni() *PizzaBuilder {
	b.pepperoni = true; return b
}

func (b *PizzaBuilder) AddVeggie(v string) *PizzaBuilder {
	b.veggies = append(b.veggies, v); return b
}

func (b *PizzaBuilder) Build() Pizza {
	return Pizza{Size: b.size, Cheese: b.cheese, Pepperoni: b.pepperoni, Veggies: b.veggies}
}

func main() {
	pizza := NewPizzaBuilder().
		SetSize("Large").
		AddCheese().
		AddPepperoni().
		AddVeggie("Mushrooms").
		Build()
	fmt.Printf("%+v\n", pizza)
}
