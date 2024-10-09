package string

import "fmt"

func main() {
	// 示例 1。
	s1 := make([]int, 5)
	fmt.Printf("The length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)

	println("=====================================")

	ia := []int{1, 2, 3, 4, 5}

	fmt.Println("slice :3 is :", ia[:3])
	fmt.Println("slice 2: is ", ia[2:])
}
