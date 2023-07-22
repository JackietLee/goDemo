package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fn1(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("写入%v成功\n", i)
		time.Sleep(time.Millisecond * 50)
	}
	close(ch)
	wg.Done()
}

func fn2(ch chan int) {
	for v := range ch {
		fmt.Printf("读取%v成功\n", v)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done()
}
func main() {

	var ch1 = make(chan int, 10)
	wg.Add(1)
	go fn1(ch1)
	wg.Add(2)
	go fn2(ch1)
	wg.Wait()
	time.Sleep(time.Millisecond * 5000)
	fmt.Println("退出")
}
