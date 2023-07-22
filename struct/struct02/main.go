package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
	sex  string
}

func (p Person) printInfo() {
	fmt.Printf("姓名：%v 年龄：%v 性别：%v\n", p.name, p.age, p.sex)
}

func (p *Person) setName() {
	p.name = "张三"
}
func main() {
	p7 := Person{
		"李四",
		10,
		"男",
	}
	p7.printInfo()
	p7.setName()
	p7.printInfo()
}
