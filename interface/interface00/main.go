package main

import "fmt"

type Usber interface {
	start()
	stop()
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
	var usb1 *Phone = new(Phone)
	usb1.Name = "手机2"
	usb.start()
	usb.stop()
	usb1.start()
	usb1.stop()

	var p1 Usber = usb
	p1.start()
	p1.stop()
}
