package main

import "fmt"

type Animal interface {
	A
	B
}

type A interface {
	GetName() string
}

type B interface {
	SetName(string2 string)
}
type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}
func (d Dog) GetName() string {
	return d.Name
}
func main() {

	d := &Dog{
		Name: "柯基",
	}
	var d1 Animal = d
	fmt.Println(d1.GetName())
	d1.SetName("哈士奇")
	fmt.Println(d1.GetName())
}
