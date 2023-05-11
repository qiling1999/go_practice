package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//web开发就是一个请求对应一个响应
func sayHello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w,"<h1>hello zhiqing</h1><h2>你好</h2>")
	//可以把上面代码中的<h1>hello zhiqing</h1><h2>你好</h2>单独放到一个文件中，例如hello中,这些是前端方面的内容
	b, _ := ioutil.ReadFile("src/practice/web/hello")
	fmt.Fprintln(w, string(b))
}

func main() {
	//后端就是指定了你访问的函数，函数控制了
	 http.HandleFunc("/hello", sayHello)

	 //启动服务
	 err := http.ListenAndServe("127.0.0.1:9090",nil)
	 if  err != nil {
		fmt.Printf("http server failed err%v", err)
		return
	 }
}
