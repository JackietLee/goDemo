package main

import "fmt"

func main() {
	var a = new(int)
	fmt.Println("值：%v 类型：%T 指针变量对应的值：%v", a, a, *a)
	var b *int
	b = new(int)
	*b = 100
	fmt.Println(*b)
}
