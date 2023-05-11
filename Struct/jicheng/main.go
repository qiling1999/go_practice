package main

import "fmt"

type A struct {
	Name string
	Age int
}

func (a *A)Sayok()  {
	fmt.Println("A ok", a.Name)
}
//B继承了匿名类A中的变量和方法
type B struct {
	A
	Score int
}

//结构体嵌入多个匿名结构体
type C struct {
	Name string
	A
	B
}

//嵌入有名结构体 称为组合
type D struct {
	a A
}
func main() {
	var b B //声明
	//匿名结构体字段有两种访问方法
	b.A.Name = "zhiqing"
	b.A.Sayok()
	b.Name = "xining"
	b.Sayok()

	//多个匿名结构体访问必须指定匿名结构体名字
	var c C
	c.A.Name = "zhiqing"
	c.B.Age = 25

	//有名结构体组合在访问变量时，必须带上结构体名字
	var d D
	d.a.Name = "zhiqing"

	//也可以在创建结构体变量时直接指定各个匿名结构体字段的值
	b1 := B{A{"zhiqing", 26}, 10000}
	b2 := B{
		A{
			Name:"zhiqing",
			Age:26,
		},
		10000,
	}
	fmt.Println("b1 b2 ",b1, b2)
}