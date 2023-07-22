package main

import "fmt"

type Usber interface {
	start()
	stop()
}
type Computer struct {
}

func (c Computer) link(u Usber) {
	u.start()
	u.stop()
}

type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "停止")
}

func main() {
	var usb Phone = Phone{
		Name: "手机1",
	}
	computer := Computer{}
	computer.link(usb)
}
