package main

import (
	"fmt"
	"math/rand"
	"time"
)

//数组相关
func main() {
	//数组定义和初始化
	//方法一：
	var numArray1 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArray1 = ", numArray1)

	//方法二
	var numArray2 = [3]int{4, 5, 6}
	fmt.Println("numArray2 = ", numArray2)

	//方法三   数组长度不固定
	var numArray3 = [...]int{7, 8, 9, 10, 11, 12}
	fmt.Println("numArray3 = ", numArray3)

	//方法四   指定下标
	var numArray4 = [...]int{1: 800, 2: 900, 0: 999}
	fmt.Println("numArray4 = ", numArray4)

	strArray5 := [...]string{1: "hello", 2: "zhiqing", 0: "xaicijian"}
	fmt.Println("strArray5 = ", strArray5)

	//数组遍历
	//方法一：常规
	for i := 0; i < len(numArray1); i++ {
		fmt.Printf("numArray1[%d] = %v\n", i, numArray1[i])
	}
	//方法二：for range
	for i, v := range numArray3 {
		//fmt.Printf("i=%v, v=%v\n", i, v)
		fmt.Printf("numArray3[%d]=%v\n", i, v)
	}

	//在其他函数中修改数组中的值
	fmt.Println("原先numArray1 = ", numArray1)
	test(&numArray1)
	fmt.Println("传址后numArray1 = ", numArray1)

	//生成内容随机的数组
	var intArr [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Intn(100)   //0-100之间的随机数
	}
	//二分法将数组倒序
	fmt.Println("交换前的数组：", intArr)
	temp := 0
	for i := 0; i < len(intArr) / 2; i++ {
		temp = intArr[len(intArr) - 1 - i]
		intArr[len(intArr) - 1 - i] = intArr[i]
		intArr[i] = temp
	}
	fmt.Println("交换后的数组：", intArr)

	//切片的使用
	//方法一，先定义一个数组，然后让切片去引用已经创建好的数组
	var arr [5]int = [...]int {1, 2, 3, 4, 5}
	var slice1 = arr[1:3]  //取出从arr[1]到第三个元素arr[2],
	//var s = arr[:] ；后面无论是空着还是写len(arr)还是写数组个数,都是切到数组的最后一个数，且包含最后一个数
	fmt.Println("slice1=", slice1)
	//方法二，通过make方法创建切片，5为len，10为cap切片容量，可选
	var slice2 []float64 = make([]float64, 5, 10)
	slice2[1] = 10
	slice2[3] = 20
	fmt.Println("slice2=", slice2)
	fmt.Println("slice2的size = ", len(slice2))
	fmt.Println("slice2的cap = ", cap(slice2))
	fmt.Println("")
	//方法三，定义一个切片，直接就指定具体数组，使用原理类似make的方式
	var Slice3 []string = []string{"haha","nihao","zhiqing"}
	fmt.Println("strSlice = ", Slice3)
	fmt.Println("strSlice size = ", len(Slice3))
	fmt.Println("strSlice cap = ", cap(Slice3))

	//切片的遍历
	//方法一，常规for循环遍历
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice1[%d] = %v\n", i, slice1[i])
	}
	//方法二：for-range结构循环遍历切片
	for i,v := range slice2 {
		fmt.Printf("i = %v v = %v\n", i, v)
	}

	//切片字符串string 并且修改字符串
	str := "nihaozhiqing"   //定义一串字符串
	slice4 := str[6:]       //切片指向字符型数组
	fmt.Println("slice4", slice4)
	arrstr1 := []byte(str)    //将字符型数组转化成byte型数组
	arrstr1[0] = 'o'          //通过[]byte更改字符串的值
	str = string(arrstr1)     //再将[]byte型数组转成字符型数组
	fmt.Println("str = ", str)
	arrstr2 := []rune(str)    //原理跟上面一样，只是转换成[]rune
	arrstr2[0] = '你'         //[]rune是按字符处理，兼容汉字
	str = string(arrstr2)
	fmt.Println("str = ", str)

	fbnSlice := test2(20)
	fmt.Println("fbnSlice = ", fbnSlice)
	fmt.Printf("fbnSlice = ", fbnSlice)

}

//数组传递地址
func test(numArray1 *[3]int) {
	(*numArray1)[0] = 0
}

//斐波那契函数
func test2 (n int) ([]uint64) {
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1

	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i - 1] + fbnSlice[i - 2]
	}
	return fbnSlice
}

