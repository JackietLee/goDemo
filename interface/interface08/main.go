package main

import "fmt"

type Usber interface {
	start()
	stop()
}

type Computer struct {
}

func (c Computer) work(usb Usber) {
	if _, ok := usb.(Phone); ok {
		usb.start()

	} else {
		usb.stop()
	}
}

type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name + "手机启动")
}
func (p Phone) stop() {
	fmt.Println(p.Name + "手机关机")
}

type Camera struct {
}

func (p Camera) start() {
	fmt.Println("相机启动")
}

func (p Camera) stop() {
	fmt.Println("相机关机")
}
func main() {
	var p1 = Phone{
		Name: "小米手机",
	}
	var p2 Usber = p1
	p2.start()

	var p3 = &Phone{
		Name: "苹果手机",
	}
	var p4 Usber = p3
	p4.start()
}
