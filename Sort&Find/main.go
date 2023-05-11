package main

import (
	"fmt"
)

//冒泡排序
func maopao(arr *[5]int) {
	fmt.Println("排序前为：", (*arr))
	temp := 0 //临时变量()用于做交换
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr) - i - 1; j++ {
			if (*arr)[j] > (*arr)[j + 1]{   //>从小到大 <从大到小
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j + 1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后为：", (*arr))
}

func main() {
	MpArr := [5]int{24, 69, 80, 57, 13}
	maopao(&MpArr)
	shunxuxhazhao()
	erfenchazhao()

}

//顺序查找
func shunxuxhazhao () {
	names := [4]string {"ni", "hao", "zhi", "qing"}
	var strName = ""
	fmt.Println("请输入要查找的字符串：")
	fmt.Scanln(&strName)

	//方法一  同时确定查找内容和下标，一起输出
	for i := 0; i < len(names); i++ {
		if strName == names[i]{
			fmt.Printf("找到要查找的字符串%v, 下标%d\n", strName, i)
			break
		}else if i == (len(names) - 1) {
			fmt.Printf("没有找到%v", strName)
		}
	}
	//方法二  先只确定查找的下标，最后输出
	index := -1
	for i := 0; i < len(names); i++ {
		if strName == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到要查找的字符串%v, 下标%d\n", strName, index)
	}else {
		fmt.Println("没有找到%v", strName)
	}
}

//二分查找   在无重复的有序数组中查找目标数下标
func erfenchazhao() {
	arr := [5]int{19, 28, 37, 46, 55}

	var num int = 0
	fmt.Println("请输入要查找的数：")
	fmt.Scanln(&num)

	low := 0
	high := len(arr) -1                                        //寻找左侧边界的二分查找  在有重复元素的排序数组中找元素的第一个的下标
	for low <= high {                                          //for low <= high{
		mid := low + (high - low) / 2                          //    mid := low + (high - low) /2
		if arr[mid] == num{                                    //    if arr[mid] == num{
			fmt.Printf("目标数字%v的下标为%d", num, mid) //         high = mid - 1
			break   //此时找到了 mid就是目标数字下标的位置        //    }else if arr[mid] > num{
		}else if arr[mid] > num {                              //         high = mid - 1
			high = mid - 1                                     //    }else {
		}else if arr[mid] < num {                              //         low = mid + 1
			low = mid + 1                                      //    }
		}                                                      //}
	}                                                          //if low >= len(arr) || arr[low] != num{
}                                                              //    return -1
                                                               //}
                                                               //return low //此时low就是所要查找的数组下标

