package main

import "fmt"

func main() {
	var a = 10
	var p = &a
	fmt.Printf("a的值：%v a的类型%T a的地址%p\n", a, a, &a)
	fmt.Printf("p的值：%v p的类型%T b的地址%p\n", p, p, &p)
	fmt.Println(*p)
	*p = 20
	fmt.Println(a)
}
