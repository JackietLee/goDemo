package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30
	a := <-ch
	fmt.Println(a)
	<-ch
	b := <-ch
	fmt.Println(b)
}
