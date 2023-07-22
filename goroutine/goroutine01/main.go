package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
