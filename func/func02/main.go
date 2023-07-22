package main

import "fmt"

type calcType func(int, int) int

func calc(x, y int, cb calcType) int {
	return cb(x, y)
}

func add(x, y int) int {
	return x + y
}

func multi(x, y int) int {
	return x * y
}
func main() {
	i := calc(10, 5, add)
	fmt.Println(i)

	j := calc(10, 5, multi)
	fmt.Println(j)
}
