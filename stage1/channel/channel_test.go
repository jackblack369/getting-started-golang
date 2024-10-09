package channel

import (
	"fmt"
	"testing"
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
