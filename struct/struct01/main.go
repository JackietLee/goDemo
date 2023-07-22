package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

func main() {
	//var p1 Person
	//p1.name = "李健"
	//p1.age = 10
	//p1.sex = "男"
	//fmt.Printf("值：%v  类型：%T\n", p1, p1)
	//fmt.Printf("值：%#v  类型：%T\n", p1, p1)
	/*	p2 := new(Person) //指针
		p2.age = 10
		p2.name = "111"
		p2.sex = "男"
		fmt.Printf("值：%v  类型：%T\n", p2, p2)
		fmt.Printf("值：%#v  类型：%T\n", p2, p2)*/
	/*	p3 := &Person{}
		p3.age = 10
		p3.name = "111"
		p3.sex = "男"
		fmt.Printf("值：%v  类型：%T\n", p3, p3)
		fmt.Printf("值：%#v  类型：%T\n", p3, p3)*/
	p4 := &Person{
		name: "李四",
		age:  10,
		sex:  "男",
	}
	fmt.Printf("值：%v  类型：%T\n", p4, p4)
	fmt.Printf("值：%#v  类型：%T\n", p4, p4)
	p5 := Person{
		name: "李四",
		age:  10,
		sex:  "男",
	}
	fmt.Printf("值：%v  类型：%T\n", p5, p5)
	fmt.Printf("值：%#v  类型：%T\n", p5, p5)
	p6 := Person{}
	p6.age = 10
	p6.name = "111"
	p6.sex = "男"
	fmt.Printf("值：%v  类型：%T\n", p6, p6)
	fmt.Printf("值：%#v  类型：%T\n", p6, p6)

	p7 := Person{
		"李四",
		10,
		"男",
	}
	fmt.Printf("值：%v  类型：%T\n", p7, p7)
	fmt.Printf("值：%#v  类型：%T\n", p7, p7)
}
