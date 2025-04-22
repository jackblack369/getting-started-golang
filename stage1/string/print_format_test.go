package string

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func Test1(t *testing.T) {
	p := Person{Name: "Alice", Age: 30}
	fmt.Printf("%v\n", p)  // Output: {Alice 30}
	fmt.Printf("%+v\n", p) // Output: {Name:Alice Age:30}
}
