package main

import "fmt"

type Animal struct {
	name string
}

func (a Animal) run() {
	fmt.Printf("%v 在运动", a.name)
}

type Dog struct {
	Age int
	Animal
}

func (d Dog) wang() {
	fmt.Printf("%v 在叫", d.name)
}
func main() {
	var d = Dog{
		Age: 10,
		Animal: Animal{
			name: "wang",
		},
	}
	d.run()
	d.wang()
}
