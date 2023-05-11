package main

import "fmt"

//二维数组相关
func main() {
	//二维数组定义和初始化
	//方法一：
	var numArray1 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("numArray1 = ", numArray1)

	//方法二
	var numArray2 [2][3]int = [...][3]int{{11, 22, 33}, {44, 55, 66}}
	fmt.Println("numArray2 = ", numArray2)

	//方法三
	var numArray3 = [2][3]int{{31, 32, 33}, {34, 35, 36}}
	fmt.Println("numArray3 = ", numArray3)

	//方法四     常用
	var numArray4 = [...][4]int{{41, 42, 43, 44}, {45, 46, 47, 48}, {49, 50, 51, 52}}
	fmt.Println("numArray4 = ", numArray4)

	//方法五
	strArray5 := [...][3]string{{"hello", "hi", "12"}, {"zhiqing", "susu", "5.19"}, {"xaicijian", "zaijian", "98"}}
	fmt.Println("strArray5 = ", strArray5)

	//二维数组遍历
	//方法一：常规for循环
	for i := 0; i < len(numArray1); i++ {
		for j := 0; j < len(numArray1[i]); j++ {
			fmt.Printf("numArray1[%d][%d] = %v ", i, j, numArray1[i][j])
		}
		fmt.Println()
	}
	//方法二：for range
	for i, v1 := range numArray2 {
		for j, v2 := range v1 {
			fmt.Printf("numArray2[%d][%d] = %v ", i, j, v2)
		}
		fmt.Println()
	}
}


