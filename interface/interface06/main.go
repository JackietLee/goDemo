package main

import "fmt"

func MyPrint1(a interface{}) {
	if _, ok := a.(string); ok {
		fmt.Println("string类型")
	} else if _, ok := a.(int); ok {
		fmt.Println("int类型")
	} else if _, ok := a.(bool); ok {
		fmt.Println("bool类型")
	}
}

// x.(type)判断一个变量的类型，只能用在switch语句里面
func MyPrint2(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case bool:
		fmt.Println("bool类型")
	default:
		fmt.Println("错误")

	}
}

func main() {
	var a interface{}
	//a = "你好golang"
	a = 11
	s, ok := a.(string)
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println("失败")
	}
}
