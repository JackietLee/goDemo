package main

import "fmt"

func main() {
	//匿名自执行函数
	func() {
		fmt.Println("test..")
	}()

	var fn = func(x, y int) int {
		return x * y
	}
	fmt.Println(fn(2, 3))

	a := func(x, y int) int {
		return x + y
	}(10, 11)
	fmt.Println(a)
}
