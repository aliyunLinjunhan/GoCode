package main

import "fmt"

// 声明/定义一个接口
type Usb interface {
	//声明了该接口的俩个方法
	Start()
	Stop()
}

type Phone struct {

}

// 让Phone 实现Usb接口的方法
func (p Phone) Start() {

	fmt.Println("手机开始工作。。。。")
}
func (p Phone) Stop() {

	fmt.Println("手机停止工作。。。。")
}
func (p Phone) Call() {

	fmt.Println("手机在打电话。。。。")
}

type Camera struct {

}

// 让Camers 实现USB接口的方法
func (p Camera) Start() {

	fmt.Println("相机开始工作。。。。")
}
func (p Camera) Stop() {

	fmt.Println("相机停止工作。。。。")
}

func main() {
	var usbSilce [3]Usb
	usbSilce[0] = Phone{}
	usbSilce[1] = Camera{}
	usbSilce[2] = Phone{}
	var computer Computer
	for _, v := range usbSilce {
		computer.Working(v)
		fmt.Println()
	}
}

// 计算机
type Computer struct {

}

// 编写一个working的方法， 接受一个Usb接口类型的变量
func (c Computer) Working(usb Usb) {

	usb.Start()
	if phone, ok := usb.(Phone); ok{
		phone.Call()
	}
	usb.Stop()
}
