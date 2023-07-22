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
}

func (p Phone) start() {
	fmt.Println("手机启动")
}
func (p Phone) stop() {
	fmt.Println("手机关机")
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
	var computer = Computer{}
	var phone = Phone{}
	var camera = Camera{}
	computer.work(phone)
	computer.work(camera)
}
