package main

import "fmt"

//多态 不同结构体实现同一个接口，使得接口变量呈现不同形态

//声明定义一个接口
type Usb interface {
	Start()
	Stop()
}

//定义一个结构体
type Phone struct {
	name string
}
//让Phone结构体实现接口Usb
func (p Phone)Start(){
	fmt.Println("手机开始工作")
}
func (p Phone)Stop(){
	fmt.Println("手机停止工作")
}
//Phone结构体自己的方法
func (p Phone)Call(){
	fmt.Println("手机打电话")
}


//定义另一个结构体让他也实现Usb接口
type Camera struct {
	name string
}
//让Camera结构体实现接口Usb
func (c Camera)Start(){
	fmt.Println("相机开始工作")
}
func (c Camera)Stop() {
	fmt.Println("相机停止工作")
}

//计算机
type Computer struct {

}
//编写一个方法Working方法，接收一个Usb接口类型变量
func (c Computer)Working(usb Usb)  { //usb变量会更具传入的实参，来判断到底是Phone还是Camera
	//通过usb接口调用Start和Stop方法
	usb.Start()
	//类型断言
	//由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言
    //如果usb是指向Phone结构体变量，则还需要调用Call方法
    if phone, ok := usb.(Phone); ok{
    	phone.Call()
	}

	usb.Stop()
}
func main() {
	//多态形式一
	//多态参数，Usb usb既可以接收手机变量，又可以接收相机变量
	computer := Computer{}
	phone := Phone{name:"zhiqing"}
	camera := Camera{"xining"}
	computer.Working(phone)
	computer.Working(camera)

	//多态形式二
	//多态数组
	//定义一个Usb接口数组，可以存放Phone和Camera的机构体变量
	var usbArr [3]Usb
	usbArr[0] = Phone{"zhiqing"}
	usbArr[1] = Phone{"susu"}
	usbArr[2] = Camera{"xining"}
	fmt.Println(usbArr)

	//遍历usbArr,如果是Phone变量，除了调用Usb接口声明的方法外，
	//还需要调用Phone特有方法Call  类型断言
	//var computer Computer
	for _, v := range usbArr{
		computer.Working(v)
		fmt.Println()
	}

}