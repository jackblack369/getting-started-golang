package _type

import (
	"fmt"
	"testing"
)

func removeItem(num int, items []int) []int {

	// remote specify num from tests
	for i, item := range items {
		if item == num {
			items = append(items[:i], items[i+1:]...)
			break
		}
	}
	return items
}

func TestRemove(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	num := 3
	items = removeItem(num, items)
	fmt.Println(items)
}
