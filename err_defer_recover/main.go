package main

import (
	"errors"
	"fmt"
	"time"
)

//可以结合单元测试实现检查错误类型
func readConf(name string)(err error) {
	if name == "config.ini" {
		//读取
		return nil
	}else {
		//返回一个自定义的错误
		return errors.New("读取文件错误。。")
	}
}
func test1() { //自定义错误类型
	err := readConf("config1.ini")
	if err != nil {
		//如果读取文件发生错误，就输出这个错误
		fmt.Println("test1程序错误", err)
		//panic(err)   //输出这个错误，并终止程序
	}
	fmt.Println("test1程序继续执行")
}

//使用defer + recover来捕获和处理异常
func sayHello(){
	for i := 0;i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("zhiqing")
	}
}
func test2(){
	defer func() {
		//捕获test2()抛出panic
		if err := recover(); err != nil{
			fmt.Println("test2发生错误", err)
		}
	}()
	//定义了一个map
	var myMap map[int]string
	myMap[0] = "知青"
}
func main() {
	go test1()  //可以结合单元测试实现检查错误类型
	go sayHello()
	go test2()  //使用defer_recover捕获协程的panic,这样即使这个协程发生错误，主线程不受影响

	for i := 0; i < 10; i++{
		fmt.Println("main() ok = ", i)
		time.Sleep(time.Second)
	}
}
