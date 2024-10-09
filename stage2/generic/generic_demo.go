package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// NewList creates and returns a new List with the given value
func NewList[T any](val T) *List[T] {
	return &List[T]{val: val}
}

// Add appends a new node with the given value to the end of the list
func (l *List[T]) Add(val T) {
	if l.next == nil {
		l.next = NewList(val)
		return
	}
	l.next.Add(val)
}

// Print displays all elements in the list
func (l *List[T]) Print() {
	fmt.Print(l.val, " ")
	if l.next != nil {
		l.next.Print()
	} else {
		fmt.Println()
	}
}

func main() {
	// Create a new list of integers
	intList := NewList(1)
	intList.Add(2)
	intList.Add(3)
	intList.Add(4)

	fmt.Println("Integer List:")
	intList.Print()

	// Create a new list of strings
	strList := NewList("Hello")
	strList.Add("World")
	strList.Add("Go")
	strList.Add("Generics")

	fmt.Println("String List:")
	strList.Print()
}
