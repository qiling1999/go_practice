package model

type Stu1 struct {
	Name string
	Age int
}
type stu2 struct {
	Name string
	Age int
}

//结构体首字母小写工厂模式, 工厂模式提供的方法首字母大写

func NewStu2(n string, a int) *stu2{
	return &stu2{
		Name : n,
		Age : a,
	}
}
