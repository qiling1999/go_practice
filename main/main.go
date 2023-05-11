package main

import (
	"errors"
	"fmt"
)

var age int = 50
var Name int = 49


func main()  {
	//var i bool = true

	//var n1 float64 = 23.6
	//var str string
	//str = fmt.Sprintf("%t",i)
	//fmt.Printf("str的类型%T, str = %q\n", str, str)
	//str = strconv.FormatFloat(n1,'f', 10, 64)
	//fmt.Printf("str的类型%T, str = %q\n", str, str)
	//var str1 string = "123154"
	//var num1 int64
	//num1, _ = strconv.ParseInt(str, 10, 64)


	//获取一个int变量num的地址，显示
	//将num的地址赋给指针ptr，并通过ptr修改num的值
	/*var num int = 10
	fmt.Printf("num的地址=%v\n", &num)


	var ptr *int = &num
	*ptr = 11
	fmt.Println("num的值变为：", num)

	fmt.Println(model.HeroName)
	var num2 float64 = 10 / 4
	fmt.Println(num2)*/
	/*var a int = 1 >> 2
	var b int = -1 >> 2
	var c int = 1 << 2
	var d int = -1 << 2*/
	/*fmt.Println(2&3)
	fmt.Println(2|3)
	fmt.Println(5|4)
	fmt.Println(-3^3)*/
	/*var score int
	fmt.Println("请输入成绩")
	fmt.Scanln(&score)

	if score == 100 {
		fmt.Println("奖励1")
	}else if score > 80 && score < 90 {
		fmt.Println("奖励2" )
	}else {
		fmt.Println("没奖励")
	}*/
	/*var a float64 = 2.0
	var b float64 = 4.0
	var c float64 = 2.0

	m := b * b - 4 * a * c
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		x2 := (-b - math.Sqrt(m)) / 2 * a
		fmt.Printf("x1 = %v\n", x1)
		fmt.Printf("x2 = %v", x2)

	}else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Printf("x1 = %v", x1)
	}else {
		fmt.Println("无解")
	}*/

	/*var x interface{}
	var y  = 10.0
	x = y

	switch i := x.(type) {
	case nil:
		fmt.Printf("i的类型,%T", i)
	case int:
		fmt.Println("2")
	case float64:
		fmt.Println("4")
	default:
		fmt.Println("3")
	}*/
	//遍历数组
	/*var str1 string = "helloworld"
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c\n", str1[i])
	}
	//遍历数组方法2
	fmt.Println()
	var str2 string = "abc~ok"
	for index, val := range str2 {
		fmt.Printf("index=%d, val=%c\n", index, val)
	}*/
	/*var classNum int = 2
	var stuNum int = 5
	var totalSum float64 = 0.0
	var passCount int = 0
	for j := 0; j <= classNum; j++ {
		sum := 0.0
		for i:= 1; i<= stuNum ; j++ {
			var score float64
			fmt.Printf("请输入第%d班 第%d个学生的成绩\n", j, i)
			fmt.Scanln(&score)
			sum += score
			if score >= 60{
				passCount++
			}
		}
	}*/
	/*var totalLevel int = 20
	for i := 1; i <= totalLevel; i++ {
		for k := 1; k <= totalLevel - i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2 * i -1; j++ {
			if j == 1 || j == 2 * i - 1 || i ==totalLevel {
				fmt.Print("*")
			}else {
				fmt.Print(" ")
			}
		}
	}
	fmt.Println()
	fmt.Println("操作")
	*/
	//test(4)
	/*num := 8
	fbn(num)
	fmt.Println("num = ", num)*/
	/*a := getSum
	fmt.Printf("a的类型%T, getSum类型为%T,\n", a, getSum)
	res := a(10, 20)
	fmt.Println("res = ", res)*/
	/*a := 10
	b := 20
	getSum(&a ,&b)
	fmt.Printf("a = %v, b = %v", a, b)*/
	/*defer fmt.Println("ok1")
	defer fmt.Println("ok2")
	res := 100
	fmt.Println("ok3", res)*/

	/*for i := 0; i <= 10; i++ {
		fmt.Println("i=", i)
	}
	var i int
	for i = 0; i <= 10; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("i的值为：", i)*/
	/*n, err := strconv.Atoi("asdada")
	if err != nil {
		fmt.Println("转换错误", err)
	}else {
		fmt.Println("转换结果为：", n)
	}*/

	/*var bytes = []byte("zhiqing")
	fmt.Printf("bytes = %v\n", bytes)*/
	/*str := "ZHIQING"
	str = strings.ToLower(str)
	//str = strings.ToUpper(str)
	fmt.Printf("str=%v\n", str)*/
	/*b := strings.HasPrefix("ftp://192.168.10.1", "hsp")
	fmt.Printf("b = %v\n", b)*/
	/*now := time.Now()
	//fmt.Printf("Month = %v\n", now.Month())
	//fmt.Printf("Month = %v\n", int(now.Month()))
	fmt.Printf("当前年月日：%d-%d-%d %d:%d:%d\n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	datestr := fmt.Sprintf("当前年月日：%d-%d-%d %d:%d:%d\n", now.Year(),
		now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Printf("datastr = %v\n", datestr)

	fmt.Printf(now.Format("2006.01.02 15:04:05"))*/

	/*start := time.Now().UnixNano()
	getSum()
	end := time.Now().UnixNano()
	fmt.Printf("执行函数的时间为：%v纳秒", end - start)*/
	test()
	fmt.Println("main()下面的代码。。。")

}
func test() {
	//使用defer + recover来捕获和处理异常
	
	err := readConf("config1.ini")
	if err != nil {
		//如果读取文件发生错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test程序继续执行")
}


func readConf(name string)(err error) {
	if name == "config.ini" {
		//读取
		return nil
	}else {
		//返回一个自定义的错误
		return errors.New("读取文件错误。。")
	}
}
/*func fbn(n int) int {
	if n == 10 {
		return 1
	}else {
		return (fbn(n + 1) + 1) * 2
 	}
}*/
/*func test(n int) {
	if n > 2 {
		n--
		test(n)
	}else {
		fmt.Println("n=",n)
	}
}*/