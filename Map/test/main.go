package main

import "fmt"

//使用map[string]map[string]string 的map类型
//key:表示用户名，是唯一的，不可以重复
//如果某个用户名存在，就将其密码修改为"888"，
//如果不存在就增加这个用户信息(包括nickname和pwd)
//编写一个函数modifyUser(user map[string]map[string]string, name string)完成上述功能

func main() {
	users := make(map[string]map[string]string, 10)
	users["zhiqing"] = make(map[string]string, 2)
	users["zhiqing"]["pwd"] = "999"
	users["zhiqing"]["nickname"] = "susu"

	modifyUser(users, "xining")
    fmt.Println(users)
}

func modifyUser(users map[string]map[string]string, name string) {
	if users[name] == nil{
		users[name] = make(map[string]string)
		users[name]["pwd"] = "888"
		users[name]["nickname"] = "昵称：" + name
	}else {
		users[name]["pwd"] = "888"
	}
}
