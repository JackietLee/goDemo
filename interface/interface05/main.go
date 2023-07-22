package main

import "fmt"

func show(a interface{}) {
	fmt.Printf("值：%v 类型：%T\n", a, a)
}

func main() {
	var m1 = make(map[string]interface{})
	m1["username"] = "张三"
	m1["age"] = 20
	m1["married"] = true

	var s1 = []interface{}{1, 2, "你好", true}
	fmt.Printf("值：%v 类型：%T\n", s1, s1)
}
