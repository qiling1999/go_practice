package main

import (
	"fmt"
	"sync"
)

//对全局变量加锁，防止出现资源争夺问题

//定义全局变量
var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//lock 是一个全局的互斥锁
	//sync 是包，synchornized同步
	//Mutex 是互斥
	lock sync.Mutex
)

func test(n int)  {
	res := 1
	for i := 1; i <= n; i++{
		res *= i
	}

	//将结果res放到myMap中，并加锁
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	//time.Sleep(time.Second * 10)
	//加锁 输出结果
	lock.Lock()
	for i, v := range myMap{
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()

}
