// Description: This code demonstrates the use of closures in Go, where a function can capture and remember the variables from its surrounding context.
// It shows how closures can maintain state across multiple calls, even when the outer function has finished executing.
// It also illustrates the concept of variable shadowing and how closures can access variables from their lexical scope.
package main

import "fmt"

const a = 10

var p = 100

func outer() func() {
	money := 65
	age := 30

	fmt.Println("Age = ", age)

	show := func() {
		money = money + a + p
		fmt.Println(money)
	}
	return show
}

func call() {
	inc1 := outer()
	inc1()
	inc1()

	inc2 := outer()
	inc2()
	inc2()
}

func main() {
	call()
}

func init() {
	fmt.Println("Halal Money")
}
