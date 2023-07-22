package main

import "fmt"

func main() {
	var a = adder()
	i := a()
	fmt.Println(i)

	var b = adder2()
	j := b(1)
	k := b(1)
	fmt.Println(j)
	fmt.Println(k)
}

func adder() func() int {
	var i = 10
	return func() int {
		return i + 1
	}
}

func adder2() func(y int) int {
	var i = 10
	return func(y int) int {
		i += y
		return i + 1
	}
}
