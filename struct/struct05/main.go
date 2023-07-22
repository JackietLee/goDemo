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
type Class struct {
	Title    string
	Students []Student
}

func main() {
	c := Class{
		Title:    "001",
		Students: make([]Student, 0),
	}
	for i := 1; i <= 10; i++ {
		s := Student{
			ID:     i,
			Gender: "男",
			Name:   fmt.Sprintf("stu_%v", i),
		}
		c.Students = append(c.Students, s)
	}
	fmt.Println(c)
	marshal, _ := json.Marshal(c)
	fmt.Println(string(marshal))
}
