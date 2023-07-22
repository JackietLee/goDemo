package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//var ch1 = make(chan int, 10)
	//for i := 1; i <= 10; i++ {
	//	ch1 <- i
	//}
	//close(ch1) // 通过forrange便利需要关闭管道
	//for i := range ch1 {
	//	fmt.Println(i)
	//}

	var ch1 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
	close(ch1) // 通过forrange便利需要关闭管道
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch1)
	}
}
