package _type

import "fmt"

// valid array types
type (
	T5 [10]*T5              // T5 contains T5 as component of a pointer
	T6 [10]func() T6        // T6 contains T6 as component of a function type
	T7 [10]struct{ f []T7 } // T7 contains T7 as component of a slice in a struct
)

// A struct with 6 fields.
type A struct {
	x, y int
	u    float32
	_    float32 // padding
	A    *[]int
	F    func()
}

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
