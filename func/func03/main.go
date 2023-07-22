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

func do(o string) calcType {
	switch o {
	case "+":
		return add
	case "*":
		return multi
	case "-":
		return func(i int, i2 int) int {
			return i - i2
		}
	default:
		return nil
	}
}
func main() {
	i := calc(10, 5, add)
	fmt.Println(i)

	j := calc(10, 5, multi)
	fmt.Println(j)

	c := do("+")
	i2 := c(1, 2)
	fmt.Println(i2)
}
