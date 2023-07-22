package main

import "fmt"

func main() {
	var arr = [...]int{12, 11, 22, 33}
	var sum = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Printf("数组的和: %v", sum)
	var slice = []int{12, 122, 22, 55, 88}
	slice = append(slice[:2], slice[3:]...)
	fmt.Println(slice)

	s2 := "你好golang"
	runeStr := []rune(s2)
	runeStr[0] = '大'
	fmt.Println(string(runeStr))
}
