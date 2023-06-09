package main

import (
	"fmt"
	"reflect"
)

//反射应用场景
//1.不知道接口调用的哪个函数，根据传入参数在运行时确定调用的具体接口，
//这种需要对函数或方法进行反射，例如桥接模式
//2.对结构体序列化时，如果结构体有指定Tag，也会使用到反射生成对应的字符串

//反射重要函数
//reflect.TypeOf(变量名) 获取变量的类型，返回reflect.Type类型
//reflect.ValueOf(变量名) 获取变量的值，返回reflect.Value类型reflect.Value是一个结构体类型

//反射最佳实例：使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值

//变量、interface{}和reflect.Value是可以相互转化的
/*
var student Stu
var num int
func test(b interface{}){
	//1.如何将interface{}转成reflect.Value
	rVal := reflect.ValueOf(b)
	//2.如何将reflect.Value转成interface{}
	iVal := rVal.interface()
	//3.如何将interface{}转成原来的变量类型，使用类型断言
	v := iVal.(Stu)
}
*/

//反射类型转换练习
type Student struct {
	Name string
	Age int
}

//演示反射类型转换
func reflectTest01(b interface{}) {

	//通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)
	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	n2 := 2 + rVal.Int()
	//n3 := rVal.Float()
	fmt.Println("n2=", n2)
	//fmt.Println("n3=", n3)
	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)
	iV := rVal.Interface()   //将reflect.Value转成interface{}
	num2 := iV.(int)         //将 interface{} 通过断言转成需要的类型
	fmt.Println("num2=", num2)
}

//演示反射类型转换[对结构体的反射]
func reflectTest02(b interface{}) {

	//通过反射获取的传入的变量的 type , kind, 值
	//1. 先获取到 reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)
	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	//3. 获取 变量对应的Kind   获取变量的类别，返回的是一个常量
	//(1) rVal.Kind() ==>
	kind1 := rVal.Kind()
	//(2) rTyp.Kind() ==>
	kind2 := rTyp.Kind()
	fmt.Printf("kind =%v kind=%v\n", kind1, kind2)
	//下面我们将 rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T \n", iV, iV)
	//将 interface{} 通过断言转成需要的类型
	//这里，我们就简单使用了一带检测的类型断言.
	//同学们可以使用 swtich 的断言形式来做的更加的灵活
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}


//反射实例
type Monster struct {
	Name string`json:"name"`
	Age int`json:"monster_age"`
	Score float32`json:"成绩"`
	Sex string
}
//结构体的方法 ，返回两个数的和
func (s Monster)GetSum(n1, n2 int) int {
	return n1 + n2
}
//结构体的方法，接收4个值，给s赋值
func (s Monster)Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
//结构体的方法，显示s的值
func (s Monster)Print(){
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}
//测试方法  测试结构体
func TestStruct(a interface{})  {
	//获取reflect.Type 类型
	typ := reflect.TypeOf(a)
	//获取reflect.Value 类型
	val := reflect.ValueOf(a)
	//获取到a对应的类别
	kd := val.Kind()
	//如果传入的不是struct，就退出
	if kd !=  reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取到该结构体有几个字段
	num := val.NumField()

	fmt.Printf("struct has %d fields\n", num) //4
	//变量结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))
		//获取到struct标签, 注意需要通过reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		//如果该字段于tag标签就显示，否则就不显示
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}

	//获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	//var params []reflect.Value
	//方法的排序默认是按照 函数名的排序（ASCII码）  所以第一个方法是Get 第二个是Print 第三个是Set
	val.Method(1).Call(nil) //获取到第二个方法。调用它


	//调用结构体的第1个方法Method(0)
	var params []reflect.Value  //声明了 []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //传入的参数是 []reflect.Value, 返回[]reflect.Value
	fmt.Println("res=", res[0].Int()) //返回结果, 返回的结果是 []reflect.Value*/
}

func main() {

	//演示对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作
	//1. 先定义一个int
	var num int = 100
	reflectTest01(num)
	//2. 定义一个Student的实例
	stu := Student{
		Name : "tom",
		Age : 20,
	}
	reflectTest02(stu)

	//反射最佳实例：使用反射来遍历结构体的字段，调用结构体的方法，并获取结构体标签的值
	//创建了一个Monster实例
	var a Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}
	//将Monster实例传递给TestStruct函数
	TestStruct(a)

}
