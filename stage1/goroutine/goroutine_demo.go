package main

import (
	"fmt"
	"sync"
)

//func main() {
//	for i := 0; i < 10; i++ {
//		go func(i int) {
//			fmt.Println(i)
//		}(i)
//	}
//	time.Sleep(1 * time.Second)
//}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1) // 增加计数器
		go func(i int) {
			defer wg.Done() // 减少计数器
			fmt.Println(i)
		}(i)
	}

	wg.Wait() // 等待所有 goroutine 完成
	fmt.Println("All goroutines finished.")
}
