package main

import (
	"fmt"
	"sync"
)

var counter = 0
var wg sync.WaitGroup
var mutex sync.Mutex

func incrementCounter() {
	mutex.Lock()
	defer mutex.Unlock()
	counter++
	wg.Done()
}

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incrementCounter()
	}

	wg.Wait()
	fmt.Println(counter)
}
