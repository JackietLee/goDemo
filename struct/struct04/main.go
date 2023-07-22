package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int    `json:"id"` // 结构体标签
	Gender string `json:"gender"`
	Name   string `json:"name"`
	Sno    string `json:"sno"`
}

func main() {
	var s1 = Student{
		ID:     12,
		Gender: "男",
		Name:   "李四",
		Sno:    "s0001",
	}
	fmt.Printf("%#v", s1)
	marshal, _ := json.Marshal(s1)
	fmt.Println(string(marshal))

	var str = `{"ID":12,"Gender":"男","Name":"李四","Sno":"s0001"}`
	var s2 Student
	json.Unmarshal([]byte(str), &s2)
	fmt.Printf("%#v", s2)
}
