package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test() {
	for i := 0; i <= 10; i++ {
		fmt.Println("test() 你好golang", i)
		time.Sleep(time.Millisecond * 5)
	}
	wg.Done()

}
func main() {
	wg.Add(1)
	go test()
	for i := 0; i <= 10; i++ {
		fmt.Println("main() 你好golang", i)
		time.Sleep(time.Millisecond * 5)
	}
	wg.Wait()
	fmt.Println("主线程退出")
}
