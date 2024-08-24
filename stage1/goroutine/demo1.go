package main

import "fmt"

type IntArray []int

func (ia IntArray) Sort() []int {
	// sort ia array
	if len(ia) < 2 {
		return ia
	}
	left, right := 0, len(ia)-1
	pivot := len(ia) / 2

	ia[pivot], ia[right] = ia[right], ia[pivot]

	for i := range ia {
		if ia[i] < ia[right] {
			ia[i], ia[left] = ia[left], ia[i]
			left++
		}
	}

	ia[left], ia[right] = ia[right], ia[left]

	IntArray(ia[:left]).Sort()
	IntArray(ia[left+1:]).Sort()

	return ia

}

func doSomethingForAWhile() {
	// do something
	println("do something for a while")
}

func main() {
	list := IntArray{1, 3, 5, 2, 4}
	c := make(chan int) // Allocate a channel.
	var res []int
	// Start the sort in a goroutine; when it completes, signal on the channel.
	go func() {
		res = list.Sort()
		// Print all items in the sorted array
		fmt.Println("sort finished:", res)
		for _, item := range res {
			fmt.Println(item)
		}
		c <- 1 // Send a signal; value does not matter.
	}()
	doSomethingForAWhile()
	<-c // Wait for sort to finish; discard sent value.

}
