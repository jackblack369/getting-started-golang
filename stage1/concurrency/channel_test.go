package concurrency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func channel1() {
	ch1 := make(chan int)
	go func() {
		ch1 <- 2
		ch1 <- 1
		ch1 <- 3
		close(ch1)
	}()
	elem1 := <-ch1
	elem2 := <-ch1
	elem3 := <-ch1

	fmt.Println("The first element:", elem1, "\nThe second element received from channel ch1:", elem2, "\nThe third element received from channel ch1:", elem3)
}

func TestCHN1(t *testing.T) {
	channel1()
}

func channel2() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	elem2 := <-ch1
	elem3 := <-ch1

	fmt.Println("The first element:", elem1, "\nThe second element received from channel ch1:", elem2, "\nThe third element received from channel ch1:", elem3)
}

func TestCHN2(t *testing.T) {
	channel2()
}

func channel3() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	// random receive item from channel
	select {
	case elem1 := <-ch1:
		fmt.Println("The first element:", elem1)
	case elem2 := <-ch1:
		fmt.Println("The second element received from channel ch1:", elem2)
	case elem3 := <-ch1:
		fmt.Println("The third element received from channel ch1:", elem3)
	default:
		fmt.Println("No element received from channel ch1")
	}

}

func TestCHN3(t *testing.T) {
	channel3()
}

func producer(ch chan int) {
	ch <- 42
	close(ch)
}

func consumer(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

func TestProCon(t *testing.T) {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func TestSum(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func receive(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
use a WaitGroup to properly synchronize the goroutines,
ensuring that the main function doesn't exit until the receiver has finished processing the channel values
*/
func TestWG(t *testing.T) {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go receive(ch, &wg)
	ch <- 1
	ch <- 2
	wg.Wait() // Wait for the receive goroutine to finish
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func TestRange(t *testing.T) {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c { // this loop receives values from the channel repeatedly until it is closed.
		fmt.Println(i)
	}
}

/*
A select blocks until one of its cases can run, then it executes that case.
It chooses one at random if multiple are ready.
*/
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for { // The for loop is essential for continuously executing the select statement, which handles sending Fibonacci numbers to the channel c and checking for a quit signal.
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func TestSelect(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

func TestDefaultSelect(t *testing.T) {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Printf("tick. %d \n", time.Now().UnixMilli())
		case <-boom:
			fmt.Printf("BOOM! %d \n", time.Now().UnixMilli())
			return
		default:
			fmt.Printf("    . %d \n", time.Now().UnixMilli())
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func TestCloseChannel(t *testing.T) {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch) // oops, don't forgot to close the channel after sending all data
	}()

	// keep receiving data from the channel
	for v := range ch {
		fmt.Println(v)
	}
}
