package parrallel

import (
	"fmt"
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2")
	}()
	wg.Wait()
	fmt.Println("All goroutines complete.")
}

func TestWaitGroup2(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go func(i int) {
			/*
				Each goroutine goes off and does its thing.
				When it’s done, it should let the WaitGroup know by calling WaitGroup.Done() to reduce the counter by one.
			*/
			wg.Add(1)
			defer wg.Done()
			fmt.Println("Task", i)
		}(i)
	}

	wg.Wait() // This pauses the main goroutine until that counter in the WaitGroup reaches zero
	fmt.Println("Done")
}
