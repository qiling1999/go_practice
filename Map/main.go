package main

import (
	"fmt"
	"sort"
)

//map映射
func main() {
	//map的定义和初始化
	//方法一
	var a map[string]string
	a = make(map[string]string, 10)
	a["角色1"] = "宋江"
	a["角色2"] = "吴用"
	a["角色3"] = "武松"
	a["角色1"] = "宋江"
	fmt.Println(a)

	//方法二     常用
	b := make(map[string]string)
	b["No1"] = "北京"
	b["No2"] = "内蒙古"
	b["No3"] = "包头"
	fmt.Println(b)

	//方法三   没用make
	c := map[int]string{
		5 : "su",
		1 : "zhi",
		9 : "qing",
	}
	c[519] = "susu"
	fmt.Println(c)

	//map增加与更新  如果已经有了Key，那就是修改，没有就是增加
	c[519] = "zhiqing"
	fmt.Println(c)

	//map查找    val为映射的值，ok 的值为true or false
	val, ok := c[5]
	if ok {
		fmt.Printf("要查找的映射的值为：%v\n", val)
	}else {
		fmt.Println("没有要查找的映射")
	}

	//map删除，没有一次删除所有map的方法，只有遍历然后逐个删除，或者map = make()，make一个新的，让原来的成为垃圾被回收
	delete(c, 519)
	fmt.Println(c)
	a = make(map[string]string)
	fmt.Println(a)

	//双层映射  类似二维数组
	studentMap := make(map[string]map[string]string) //给总体映射分配内存
	//给每一个键对应的映射分配内存，不能少这一句，每一个键都得写这一句
	studentMap["stu01"] = make(map[string]string)
	studentMap["stu01"]["name"] = "susu"
	studentMap["stu01"]["sex"] = "女"
	studentMap["stu01"]["address"] = "内蒙"
	studentMap["stu02"] = make(map[string]string)
	studentMap["stu02"]["name"] = "xining"
	studentMap["stu02"]["sex"] = "男"
	studentMap["stu02"]["address"] = "青岛"
	fmt.Println(studentMap)
	fmt.Println(studentMap["stu01"])
	fmt.Println(studentMap["stu01"]["name"])

	//map遍历   使用for-range遍历
	for k, v := range a{
		fmt.Printf("mapa中键%v对应的值为%v\n", k, v)
	}
	//二维映射的遍历即：map[string]map[string]string
	for k1, v1 := range studentMap{
		fmt.Println("k1", k1)
		for k2, v2 := range v1{
			fmt.Printf("k2=%v, v2=%v \n", k2, v2)
		}
	}

	//map的长度 len(studentMap)
	fmt.Println(len(studentMap))  //输出为2

	//map的切片，将数组的数据类型定义为map[string]string
	var mon []map[string]string
	mon = make([]map[string]string, 2)
	if mon[0] == nil{
		mon[0] = make(map[string]string, 2)
		mon[0]["name"] = "suzhiqing"
		mon[0]["age"] = "26"
	}
	if mon[1] == nil {
		mon[1] = make(map[string]string)
		mon[1]["name"] = "xining"
		mon[1]["age"] = "25"
	}
	newMon := map[string]string{
		"name" : "nobody",
		"age" : "0",
	}
	mon = append(mon, newMon)
	fmt.Println(mon)

	//map排序  是对key进行排序
	fmt.Println(c)
	var keys []int
	for k, _ := range c{
		keys = append(keys, k)
	}
	sort.Ints(keys)   //排序
	fmt.Println(keys)
	for _, k := range keys{
		fmt.Printf("c[%v]=%v\n", k, c[k])
	}

	//map的传址
	test(b)
	fmt.Println(b)

	//map中的value为struct的情况
	stu := make(map[string]Stu, 10)
	stu1 := Stu{"zhiqing", 26, "内蒙"}
	stu2 := Stu{"xining", 25, "青岛"}
	stu["No1"] = stu1
	stu["No2"] = stu2
	fmt.Println(stu)
	for k, v := range stu{
		fmt.Printf("编号为:%v\n", k)
		fmt.Printf("姓名为:%v\n", v.Name)
		fmt.Printf("年龄为:%v\n", v.Age)
		fmt.Printf("地址为:%v\n", v.Address)
	}
}

//定义一个学生结构体
type Stu struct {
	Name string
	Age int
	Address string
}

//map当做函数参数进行传址
func test(b map[string]string) {
	b["No3"] = "青岛"
}
