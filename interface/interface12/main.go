package main

import "fmt"

type Address struct {
	Name  string
	Phone int
}

func main() {

	var userInfo = make(map[string]interface{})
	userInfo["username"] = "张三"
	userInfo["age"] = 20
	userInfo["hobby"] = []string{"睡觉", "吃饭"}
	fmt.Println(userInfo["age"])
	fmt.Println(userInfo["hobby"])
	//fmt.Println(userInfo["hobby"][1])
	var address = Address{
		Name:  "李四",
		Phone: 12222222,
	}
	fmt.Println(address.Name)
	userInfo["address"] = address
	fmt.Println(userInfo["address"])
	strings := userInfo["hobby"].([]string)
	fmt.Println(strings[1])
	a := userInfo["address"].(Address)
	fmt.Println(a.Name)
}
