package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//实现对Hero结构体切片的排序，sort.Sort(data Interface)

//声明Hero结构体
type Hero struct {
	Name string
	Age int
}
//声明一个Hero结构体切片类型
type HeroSlice []Hero

//实现Interface接口
func (hs HeroSlice)Len() int{
	return len(hs)
}

///Less方法就是决定你用什么标准进行排序
//1、按照Hero的年龄大小进行排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice)Swap(i,j int) {
	//交换
	//temp := hs[i]
	//hs[i] = hs[j]
	//hs[j] = temp
	//上面三句话等于下面这一句话
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄|%d",rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		//将hero append到heroes切片
		heroes = append(heroes, hero)
	}

	//看看排序前的顺序
	for _, v := range heroes{
		fmt.Println(v)
	}

	//调用sort.Sort
	sort.Sort(heroes)
	fmt.Println("===排序后===")
	for _, v := range heroes{
		fmt.Println(v)
	}
}