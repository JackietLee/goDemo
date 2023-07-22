package main

import "fmt"

type Animal interface {
	SetName(string)
	GetName() string
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

type Cat struct {
	Name string
}

func (c *Cat) SetName(name string) {
	c.Name = name
}
func (c Cat) GetName() string {
	return c.Name
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
