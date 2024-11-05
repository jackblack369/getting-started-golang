package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// Function that simulates some work
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate a time-consuming task
	fmt.Printf("Worker %d done\n", id)
}

func TestWorker(t *testing.T) {
	// Set the number of OS threads the Go scheduler can use
	runtime.GOMAXPROCS(1)

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Starting 10,000 goroutines
	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id)
		}(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	fmt.Println("All workers completed")
}
