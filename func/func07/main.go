package main

import "fmt"

func f1() {
	fmt.Println("开始")
	defer func() {
		fmt.Println("aaa")
		fmt.Println("bbb")
	}()
	fmt.Println("结束")
}

func f2() int {
	var a int
	defer func() {
		a++
	}()
	return a
}

func f3() (a int) {
	defer func() {
		a++
	}()
	return a
}
func main() {
	fmt.Println(f2())
	fmt.Println(f3())
}
