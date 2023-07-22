package main

import "fmt"

func main() {
	username := "张三"
	fmt.Println(username)

	//const (
	//	n1 = iota
	//	n2
	//	n3
	//)
	//const (
	//	n1 = iota
	//	n2 = 100
	//	n3 = iota
	//)
	const (
		n1, n2 = iota + 1, iota + 2
		n3, n4
		n5, n6
	)
}
