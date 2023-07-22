package main

import "fmt"

func sumFunc(x int, y int) int {
	sum := x + y
	return sum
}

func main() {
	sum1 := sumFunc(12, 3)
	fmt.Println(sum1)
}

func sum1(x ...int) int {
	sum := 0
	for _, i := range x {
		sum += i
	}
	return sum
}

func calc(x, y int) (int, int) {
	return x + y, x - y
}
func calc1(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}

func sortIntAsc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return slice
}

func sortIntAsc1(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				temp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = temp
			}
		}
	}
	return slice
}
