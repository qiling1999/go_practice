package main

import (
	"encoding/json"
	"fmt"
	"practice/Struct/model"
)

type Person struct {
	Name string
	Age int
	Hobby string
	Score [5]float64
	ptr *int//指针
	slice []int//切片
	map1 map[string]string//map
}
type Student struct {
	Name string
	Age int
}
//树根节点，叶子节点
type Point struct {
	x int
	y int
}
type Rect struct {
	leftUp, rightDown Point
}
type Rect2 struct {
	leftUp, rightDown *Point
}

type Monster struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func (p Person) test()  {
	fmt.Println("test() name=", p.Name)
}

func main() {
	//结构体变量直接声明方法一
	var p1 Person
	p1.Name = "xining"
	p1.Age = 25
	p1.Hobby = "吃饭"

	fmt.Println("信息如下：")
	fmt.Println("Name=", p1.Name)
	fmt.Println("Age=", p1.Age)
	fmt.Println("Hobby", p1.Hobby)
	if p1.ptr == nil {
		fmt.Println("ok1")
	}
	if p1.slice == nil {
		fmt.Println("ok2")
	}
	if p1.map1 == nil{
		fmt.Println("ok3")
	}
	p1.slice = make([]int, 10)
	p1.slice[0] = 100
	p1.map1 = make(map[string]string)
	p1.map1["key1_1"] = "susu"
	fmt.Println("p1=",p1)
	p1.test()//调用方法

	//结构体变量直接声明方法二
	//p2 := Person{"susu", 26, "吃饭", nil, nil, nil, map[string]string{}}
	//fmt.Println(p2)

	//结构体变量直接声明方法三
	var p3 *Person = new(Person)
	(*p3).Name = "zhiqing"
	p3.Age = 27
	p3.slice = make([]int, 10)    //  尽管分配过值了，但是存储在不同的内存空间中。
	p3.slice[0] = 200             //  所以p1和p3不同的值，
	p3.map1 = make(map[string]string) //注意，p1和p3要分别用make创建空间
	p3.map1["key3_1"] = "daoer"
	fmt.Println("p3=",p3)  //p3= &{}
	fmt.Println("p1=",p1)

	//结构体变量直接声明方法四
	var p4 *Person = &Person{}     //p4为一个没名字的结构体变量的地址，指向的是一个没有变量名的结构体变量的0.
	
	(*p4).Name = "tom"
	p4.Name = "scout"    //重新给p4结构体变量赋值
	fmt.Println("p4=", p4)   //p4= &{}

	//var p5 *Person = &p1
	var p5 *Person = p4    //因为p4= &{}
	fmt.Println(p5.Age)
	fmt.Println((*p5).Age)
	p5.Name = "qiling"
	fmt.Printf("p5.Name=%v, p1.Name=%v\n", p5.Name, p4.Name)
	fmt.Printf("p1的地址%p\n", p4)
	fmt.Printf( "p5的地址%p, p5的值%p\n", &p5, p5)

	tree()
	json_xuliehua()
	//n1 := 10
	//n2 := 20
	p1.Age = 30
	res := p1.personTest(p1.Age, p1.Age)
	fmt.Println("res = ",res)
	res2 := p1.personTest_1()
	fmt.Println("res = ",res2)

	//如果一个类型实现了String()方法，那么fmt.Println默认会调用这个String()方法进行输出
	fmt.Println(&p1)

	//在创建结构提变量时，直接指定字段值
	var st1 = Student{"zhiqing", 26}
	st2 := Student{"zhiqing", 26}
	var st3 = Student{
		Name: "zhiqing",
		Age:  26,
	}
	st4 := Student{
		Name: "zhiqing",
		Age:  26,
	}
	//返回结构体的指针类型
	var st5 *Student = &Student{"zhiqing", 26}
	st6 := &Student{"zhiqing", 26}
	fmt.Println(st1,st2,st3,st4,st5,st6)

	//调用其他包的结构体
	//其他包结构体首字母大写，直接导包调用
	var st7 = model.Stu1{
		Name: "xining",
		Age:  25,
	}
	fmt.Println(st7)
	//其他包结构体首字母小写，用工厂模式
	var st8 = model.NewStu2("zhiqing", 26)
	fmt.Println(*st8)

}

func tree() {
	r1 := Rect{Point{1,2}, Point{3,4}}
	fmt.Printf("r1.leftUp.x的地址=%p r1.leftUp.y的地址=%p r1.rightDown.x的地址=%p r1.rightDown.y的地址=%p\n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)

	//r2有两个 *Point类型，这个两个*Point类型的本身地址也是连续的。
	//但是他们指向的地址不一定是连续的
	r2 := Rect2{&Point{10,20}, &Point{30,40}}
	fmt.Printf("r2.leftUp本身的地址=%p r2.rightDown本身的地址=%p\n",&r2.leftUp, &r2.rightDown)

	fmt.Printf("r2.leftUp指向的地址=%p r2.rightDown指向的地址=%p\n",r2.leftUp, r2.rightDown)
}
func json_xuliehua() {
	//1、创建一个Monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇"}
	//2、将monster变量序列化为json格式字串
	// json.Marshal 函数中使用反射
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json处理错误", err)
	}
	fmt.Println("jsonStr", string(jsonStr))
}
//方法不改变结构体的值
func (p1 Person) personTest(n1 int, n2 int) int {
	fmt.Println("person_test=", p1.Name)
	return n1 + n2  //p1.Age = 30
}
//方法改变结构体的值
func (p1 *Person) personTest_1() int {
	fmt.Println("person_test=", p1.Name)
	p1.Age = 40
	return p1.Age + p1.Age  //p1.Age = 40
}
//如果一个类型实现了String()方法，那么fmt.Println默认会调用这个String()方法进行输出
func (p1 *Person) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", p1.Name, p1.Age)
	return str
}






