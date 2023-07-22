package main

import "fmt"

type A interface {
}

func main() {
	var a interface{}
	a = 20
	fmt.Printf("值：%v 类型：%T\n", a, a)
	a = "你好golang"
	fmt.Printf("值：%v 类型：%T\n", a, a)
	a = true
	fmt.Printf("值：%v 类型：%T\n", a, a)
}
