package main

import "fmt"

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

func channel4() {
	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	close(ch1)

	for {
		select {
		case elem, ok := <-ch1:
			if !ok {
				return // 通道关闭，退出循环
			}
			fmt.Println("Received element:", elem)
		default:
			fmt.Println("No element received from channel ch1")
		}
	}
}

func main() {
	println("=====channel1=====")
	channel1()
	println("=====channel2=====")
	channel2()
	println("=====channel3=====")
	channel3()
	println("=====channel4=====")
	channel4()
}
