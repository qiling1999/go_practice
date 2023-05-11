package main
//gin框架是舞台，你自己的是才艺，gin是一套web框架，无重复造轮子
//关于web，web是基于HTTP协议进行交互的应用网络，通过使用浏览器/APP访问的各种资源
//后端的内容，gin框架是基于httprouter这个库开发的web框架，速度很快

import "github.com/gin-gonic/gin"


func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	//GET:请求方式； /hello:请求路径
	//当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context){
		//c.JSON:返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "hello world!",
		})
	})
	//启动HTTP服务，默认在0.0.0.0：8080启动服务
	r.Run()
}