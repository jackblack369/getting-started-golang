package stage1

import "fmt"

func main() {
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i
			//bitwise OR operation, equivalent to numbers1[i] = numbers1[i] | i, which is 4 | 3 = 7
		}
	}
	fmt.Println(numbers1)
}

func test1() {
	value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value1[1] {
	case value1[0], value1[1]:
		fmt.Println("0 or 1")
	case value1[2], value1[3]:
		fmt.Println("2 or 3")
	case value1[4], value1[5], value1[6]:
		fmt.Println("4 or 5 or 6")
	}
}
