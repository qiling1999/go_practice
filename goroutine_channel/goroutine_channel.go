package main

import "fmt"

//开启一个writeData协程，向管道intChan中写入50个整数
//开启一个readData协程，从管道intChan中读取writeData写入的数据
//注意：writeData和readData使用的是同一个管道
//主线程需要等待writeData和readData协程都完成工作之后才能退出管道

func writeData(intChan chan int){
	fmt.Printf("writeData ")
	for i := 0; i < 50; i++ {
		//放入数据
		intChan <- i
		fmt.Printf("", i)
	}
	fmt.Println("")
	close(intChan)
}
func readData(intChan chan int, exitChan chan bool){
	for{
		v, ok := <- intChan
		fmt.Printf("ok=%v\n", ok)
		if !ok {
			break
		}
		fmt.Printf("readData读取到数据=%v\n", v)
	}


	close(exitChan)
}

func main() {
	//创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		v, ok := <- exitChan
		fmt.Println("v",v,"ok",ok)
		if !ok {
			break
		}
	}
}


