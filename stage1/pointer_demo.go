package stage1

import "fmt"

// &x gets the address of x and assigns it to p.
// *p dereferences p, allowing you to read or modify the value of x through the pointer.
func main() {
	var x int = 10
	var p *int = &x // & gets the address of x

	fmt.Println("Value of x:", x)   // Outputs: Value of x: 10
	fmt.Println("Address of x:", p) // Outputs: Address of x: 0x...

	*p = 20                           // * dereferences p to change the value at the address it points to
	fmt.Println("New value of x:", x) // Outputs: New value of x: 20
}
