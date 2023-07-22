package main

import (
	"fmt"
	"strings"
)

func main() {
	var userinfo = make(map[string]string)
	userinfo["username"] = "张三"
	userinfo["age"] = "20"
	userinfo["sex"] = "男"
	var userinfo1 = map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
	}
	fmt.Println(userinfo1["username"])
	for k, v := range userinfo1 {
		fmt.Printf("key:%v value:%v\t", k, v)
	}

	var userinfo2 = make([]map[string]string, 3, 3)
	if userinfo2[0] == nil {
		userinfo2[0] = make(map[string]string)
		userinfo2[0]["username"] = "张三"
		userinfo2[0]["age"] = "20"
	}
	if userinfo2[1] == nil {
		userinfo2[1] = make(map[string]string)
		userinfo2[1]["username"] = "张三"
		userinfo2[1]["age"] = "20"
	}
	fmt.Println(userinfo2)
	fmt.Println(userinfo2[2] == nil)
	fmt.Println(userinfo2[2])
	for _, v := range userinfo2 {
		for s, s2 := range v {
			fmt.Printf("key:%v   value:%v", s, s2)
		}
	}

	var str = "    how do you do ? how     "
	var strSlice = strings.Split(strings.TrimSpace(str), " ")
	fmt.Println(strSlice)
	var strMap = make(map[string]int)
	for _, v := range strSlice {
		strMap[v]++
	}
	fmt.Println(strMap)
}
