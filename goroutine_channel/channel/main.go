package main

import (
	"fmt"
	"time"
)

type Stu struct {
	Name string
	Age int
}

//ch chan<- int 管道只读
func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++{
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}

//ch <-chan int 管道只写
func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		v, ok := <- ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan <- a
}
func main() {
	////管道可以设置为只读或者只写
	var chan1 <-chan int //只读  这里因为管道中没有东西，所以输出不了
	chan1 = make(chan int, 1)
	numchan1 := <- chan1
	fmt.Println("numchan1=", numchan1)

	var chan2 chan<- int//只写
	chan2 = make(chan int, 2)
	chan2 <- 10
	fmt.Println("chan2", chan2)

	//管道只读只写实例
	var ch chan int
	ch = make(chan int, 10)
	exitChan := make(chan struct{}, 2)
	go send(ch, exitChan)
	go recv(ch, exitChan)
	var total = 0
	for _ = range exitChan{
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("结束")


	//int类型的管道  写入读取演示
	var intChan chan int
	intChan = make(chan int, 3)
	intChan <- 10//3.向管道中写入数据
	num := 211
	intChan <- num
	intChan <- 50
	var num2 int//5.从管道中读取数据
	num2 =  <- intChan      //取出了一个，不是全部取出
	fmt.Println("num2=", num2)
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))

	//map型管道写入读取演示
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m1["city1"] = "北京"
	m1["city2"] = "内蒙"
	m2 := make(map[string]string, 20)
	m2["name1"] = "xining"
	m2["name2"] = "zhiqing"
	mapChan <- m1
	mapChan <- m2
	//m3 := make(map[string]string, 20)
	//m3 = <- mapChan
	m3 := <- mapChan  //可以不用提前定义一个空的 := 会自动判断数据类型
	fmt.Println("m3=", m3)

	//结构体变量stuChan
	var stuChan chan Stu
	stuChan = make(chan Stu ,10)
	stu1 := Stu{Name:"zhiqing", Age:26}
	stu2 := Stu{"xining", 25}
	stuChan <- stu1
	stuChan <- stu2
	stu3 := <- stuChan
	fmt.Println("stu3", stu3)

	//结构体指针变量*Stu
	var StuChan chan *Stu
	StuChan = make(chan *Stu, 10)
	stu4 := Stu{"zhiqing", 26}
	stu5 := Stu{"xining", 25}
	StuChan <- &stu4
	StuChan <- &stu5
	stu6 := <- StuChan
	fmt.Println("stu6=", stu6)

	//空接口类型，可以存放任意数据类型
	var allChan chan interface{}
	allChan = make(chan interface{}, 5)
	stu7 := Stu{"zhiqing", 26}
	var num1 = 10.2
	allChan <- stu7
	allChan <- num1
	allChan <- "susu"
	v1 := <- allChan
	v2 := <- allChan
	v3 := <- allChan
	fmt.Printf("v1=%v, v2=%v, v3=%v", v1, v2 ,v3)
	v0 := v1.(Stu) //使用类型断言  不能直接输出v1.Name
	// 由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言
	fmt.Println("v1=", v0.Name)


	//channel的关闭， close()关闭后不能写入但能读取
	int1Chan := make(chan int, 10)
	int1Chan <- 100
	int1Chan <- 200
	close(int1Chan)
	i1 := <- int1Chan
	fmt.Println("i1=", i1)


	//channel遍历for range 必须关闭管道  且不能使用普通for循环遍历
	int2Chan := make(chan int, 10)
	for i := 0; i < 10; i++{
		int2Chan <- i * i
	}
	close(int2Chan) //必须关闭管道
	for v := range int2Chan{
		fmt.Println("v= ", v)
	}

	//使用select解决从管道中读取数据的阻塞问题，因为在开发中不好确定什么时候关闭管道
	intChan10 := make(chan int, 10)
	for i := 0; i < 10; i++{
		intChan10 <- i
	}
	strChan5 := make(chan string, 5)
	for i := 0; i < 5; i++{
		strChan5 <- "hello" + fmt.Sprintf("%d",i)
	}
	for {
		select {
		case v := <- intChan10:
			fmt.Printf("从intChan10中读取到数据%d", v)
			time.Sleep(time.Second)
		case v := <- strChan5:
			fmt.Printf("从strChan5中读取到数据%s", v)
			time.Sleep(time.Second)
		default:
			fmt.Printf("都读取不到了，")
			time.Sleep(time.Second)
			return
		}
	}
}
