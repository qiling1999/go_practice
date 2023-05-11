package model

import "fmt"

type person struct {
	Name string
	age int
}

//工厂模式   只能看到姓名
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//为了访问age和sal，编写Set和Get方法
func (p *person)SetAge(age int)  {
	if age > 0 && age < 150{
		p.age = age
	}else {
		fmt.Println("输入的年龄范围错误")
	}
}

func (p *person) GetAge() int {
	return p.age
}