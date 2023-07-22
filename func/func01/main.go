package main

import "fmt"

type calc func(int, int) int

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y

}

type myInt int

func main() {
	var c calc
	c = add
	i := c(1, 2)
	fmt.Printf("c的类型：%T", c)
	fmt.Println()
	fmt.Printf("i的值：%v", i)
	var b myInt = 20
	var a int = 10
	fmt.Printf("a的类型：%T b的类型：%T\n", a, b)
	fmt.Println(a + int(b))
	fmt.Println(myInt(a) + b)
}
