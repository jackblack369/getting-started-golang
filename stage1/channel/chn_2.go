package main

import (
	"fmt"
)

func producer(ch chan int) {
	ch <- 42
	close(ch)
}

func consumer(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}
