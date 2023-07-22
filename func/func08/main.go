package main

import "fmt"

func f1() {
	fmt.Println("fn1")
}

func f2() int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("获得一个异常")
		}

	}()
	panic("异常")
}

func f3(a, b int) int {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error", err)
		}
	}()
	return a / b
}
func main() {
	fmt.Println(f2())
	i := f3(10, 0)
	fmt.Println(i)
	fmt.Println("结束")
}
