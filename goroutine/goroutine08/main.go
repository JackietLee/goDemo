package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	//使用select来获取channel里面的数据不需要关闭channel
	for true {
		select {
		case v := <-intChan:
			fmt.Printf("从 intChan 读取的数据%的\n", v)
		case v := <-stringChan:
			fmt.Printf("从 stringChan 读取的数据%的\n", v)
		default:
			fmt.Printf("数据获取完毕")
			return
		}
	}
}
