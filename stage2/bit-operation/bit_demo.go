package main

import "fmt"

func main() {
	/*
		    5: 101
			9:1001
	*/
	fmt.Println("Bitwise AND Operator 5&9:", 5&9) // 5&9 = 1
	fmt.Println("Bitwise OR Operator 5|9:", 5|9)  // 5|9 = 13
	/*
		异或运算满足交换律和结合律: a XOR b = b XOR a，a XOR (b XOR c) = (a XOR b) XOR c

		任何数和0做异或运算，结果仍然是原数：a XOR 0 = a。
		任何数和自身做异或运算，结果是0：a XOR a = 0。
		不同的数做异或运算，结果是它们之间的差异。
	*/
	fmt.Println("Bitwise XOR Operator 5^9:", 5^9)       // 5^9 = 12
	fmt.Println("Bitwise AND NOT Operator 5&^9:", 5&^9) // 5&^9 = 5
}
