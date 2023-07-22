package main

import "fmt"

func main() {
	i := fn2(10)
	fmt.Println(i)

	j := fn3(5)
	fmt.Println(j)
}

func fn1(n int) {
	if n > 0 {
		fmt.Println(n)
		n--
		fn1(n)
	}
}
func fn2(n int) int {
	if n < 0 {
		return 0
	}
	return n + fn2(n-1)
}
func fn3(n int) int {
	if n < 1 {
		return 1
	}
	return n * fn3(n-1)
}
