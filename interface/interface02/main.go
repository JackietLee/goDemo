package main

import "fmt"

type A interface {
}

func main() {
	var a A
	var str = "你好golang"
	a = str
	fmt.Printf("值：%v 类型：%T", a, a)
	var num = 20
	a = num
	fmt.Printf("值：%v 类型：%T", a, a)
	var flag = true
	a = flag
	fmt.Printf("值：%v 类型：%T", a, a)
}
