package _func

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	value, ok := interface{}(container).(map[int]string) // assert
	//fmt.Println("container is []string:", ok, " and value is ", value)
	fmt.Println("container is map[int]string:", ok, " and value is ", value)

	fmt.Printf("The element is %q.\n", container[1])

	res := string(-1)
	fmt.Println("res is ", res)

}
