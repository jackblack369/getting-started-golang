package main

import "fmt"

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
	name     string
	category string
}

func (a Dog) SetName(name string) {
	fmt.Println("Dog SetName")
	a.name = name
}

func (a Dog) Name() string {
	fmt.Println("Dog Name")
	return a.name
}

func (a Dog) Category() string {
	fmt.Println("Dog Category")
	return a.category
}

func main() {
	var dog1 *Dog
	fmt.Println("The first dog is nil. [wrap1]")
	dog2 := dog1
	fmt.Println("The second dog is nil. [wrap1]")
	var pet Pet = dog2
	if pet == nil {
		fmt.Println("The pet is nil. [wrap1]")
	} else {
		fmt.Println("The pet is not nil. [wrap1]")
	}
}
