package main

import "fmt"

type Animal interface {
	scream()
}

type Dog struct{}

func (d Dog) scream() {
	fmt.Println("Bark!!")
}

type Cat struct{}

func (c Cat) scream() {
	fmt.Println("Meow!!")
}

func main() {
	var dog Animal = Dog{}
	var cat Animal = Cat{}

	dog.scream()
	cat.scream()
}