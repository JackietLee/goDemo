package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test() {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		fmt.Println("test() 你好golang", i)
		time.Sleep(time.Millisecond * 5)
	}
}
func main() {
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()
	fmt.Println("结束")
}
