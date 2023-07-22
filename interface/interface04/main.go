package main

import "fmt"

func show(a interface{}) {
	fmt.Printf("值：%v 类型：%T\n", a, a)
}

func main() {
	show(20)
	show("你好golang")
	slice := []int{1, 2, 34, 56}
	show(slice)
}
