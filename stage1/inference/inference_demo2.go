package main

import "fmt"

type Animal interface {
	ScientificName() string
	Category() string
}

type Pet2 interface {
	Animal
	Name() string
}

type Dog2 struct {
	name string
}

func (dog Dog2) ScientificName() string {
	return "Dog"
}

func (dog Dog2) Category() string {
	return "Animal"
}

func (dog Dog2) Name() string {
	return dog.name
}

func main() {
	var dog1 *Dog2
	fmt.Println("The first dog is nil. [wrap1]")
	dog2 := dog1
	fmt.Println("The second dog is nil. [wrap1]")
	var pet Pet2 = dog2
	if pet == nil {
		fmt.Println("The pet is nil. [wrap1]")
	} else {
		fmt.Println("The pet is not nil. [wrap1]")
	}
}
