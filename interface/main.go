package main

import "fmt"

//定义一个接口，接口中有未实现的方法，  接口中不能有任何变量
//只有将所有方法都实现才可以将实例变量赋值给接口
type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}

//接口也可继承多个别的接口  ,如果要实现接口C，需要将A和B中的方法也都实现
type CInterfaca interface {
	AInterface
	BInterface
	See()
}

//定义一个结构体
type Stu struct {
	Name string
}
//创建结构体变量，并让该结构体重写接口方法
func (stu Stu) Say(){
	fmt.Println("Stu Say()")
}

//自定义数据类型也可以实现接口，
type integer int
func (i integer)Say() {
	fmt.Println("integer Say i = ", i)
}

//一个结构体实现多个接口
type Monster struct {

}
func (m Monster)Say()  {
	fmt.Println("Monster Say()~~")
}
func (m *Monster)Hello()  {     //注意这里是指针的形式，下面要对应
	fmt.Println("Monster Hello()~~")
}


func main() {
	var stu Stu //结构体变量，实现了Say()方法，实现了AInterface
	var a AInterface = stu
	a.Say()

	var i integer = 10//自定义数据类型实现Say()方法，实现了AInterface
	var b AInterface = i //给接口赋值
	b.Say()

	var monster Monster//结构体实现多个接口
	var a2 AInterface = monster
	var b2 BInterface = &monster    //因为上面是指针，下面要对应
	a2.Say()
	b2.Hello()


}
