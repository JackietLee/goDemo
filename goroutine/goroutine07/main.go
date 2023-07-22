package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	m1 := <-ch1
	fmt.Println(m1)
	ch2 := make(chan<- int, 2)
	ch2 <- 10
	ch3 := make(<-chan int, 2)
	//ch3 <- 23
	m2 := <-ch3
	fmt.Println(m2)
}
