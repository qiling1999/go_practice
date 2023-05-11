package main

import (
	"fmt"
	"redigo/redis"
)

//定义一个全局的pool
var pool *redis.Pool
//当启动程序时就初始化连接池
func init(){
	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {//初始化链接的代码，链接那个ip的redis
			return redis.Dial("tcp", "localhost:6379")
		},
		DialContext:     nil,
		TestOnBorrow:    nil,
		MaxIdle:         8,//最大空闲链接数
		MaxActive:       0,//表示和数据库的最大链接数，0表示没有限制
		IdleTimeout:     100,//最大空闲时间
		Wait:            false,
		MaxConnLifetime: 0,
	}
}

func main() {
	//通过go向redis中写入数据和读取数据
	//1.链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil{
		fmt.Println("redis.Dial err = ", err)
		return
	}
	defer conn.Close()

	//String[key-val]一个一个写入数据
	_, err = conn.Do("Set", "name", "zhiqing")
	if err != nil{
		fmt.Println("set err = ", err)
		return
	}
	//String[key-value]一个一个读取数据
	r1, err := redis.String(conn.Do("Get", "name"))
	if err != nil{
		fmt.Println("set err =", err)
		return
	}
	fmt.Println("操作ok", r1)

	//String[key-val]批量写入数据
	_, err = conn.Do("MSet", "name", "zhiqing", "age", "26")
	if err != nil{
		fmt.Println("mset err = ", err)
		return
	}
	//String[key-value]批量读取数据
	r2, err := redis.Strings(conn.Do("MGet", "name", "age"))
	if err != nil{
		fmt.Println("mset err =", err)
		return
	}
	for _, v := range r2 {
		fmt.Println("r2",  v)
	}


	//Hash写入数据（一个一个写入）
	_, err = conn.Do("HSet", "user01", "name", "zhiqing")
	if err != nil{
		fmt.Println("hset err = ", err)
		return
	}
	_, err = conn.Do("HSet","user01", "age", "26")
	if err != nil{
		fmt.Println("hset err = ", err)
		return
	}
	//Hash读取数据（一个一个读取）
	r3, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil{
		fmt.Println("hget err = ", err)
		return
	}
	r4, err := redis.String(conn.Do("HGet", "user01", "name"))
	if err != nil{
		fmt.Println("hget err = ", err)
		return
	}
	fmt.Printf("操作ok r1=%v r2=%v\n", r3, r4)

	//Hash写入数据（批量写入）
	_, err = conn.Do("HMSet", "user02", "name", "zhiqing", "age", "26")
	if err != nil{
		fmt.Println("hmset err = ", err)
		return
	}
	//Hash读取数据（批量读取）
	r5, err := redis.Strings(conn.Do("HMGet", "user02", "name", "age"))
	if err != nil{
		fmt.Println("hmget err = ", err)
		return
	}
	for i, v := range r5 {
		fmt.Printf("r[%d]=%s\n", i, v)
	}

	//给name数据设置有效时间为10s
	//_, err = conn.Do("expire", "name", 10)

/*  //List列表写入和读取数据核心代码
	_, err = conn.Do("lpush", "heroList", "no1:zhiqing", 26, "no2:xining", 25)
	r, err := redis.String(conn.Do("rpop", "heroList"))
*/

	//先从pool取出一个链接
	connPool1 := pool.Get()
	defer connPool1.Close()

	_, err = connPool1.Do("Set", "namePool", "zhiqing")
	if err != nil {
		fmt.Println("connPool.Do err = ",err)
		return
	}
	//取出
	rp1, err := redis.String(connPool1.Do("Get", "namePool"))
	if err != nil {
		fmt.Println("connPool1.Do err= ", err)
		return
	}
	fmt.Println("rp1= ", rp1)

	//如果我们要从pool链接池中取出链接，一定保证连接池是没有关闭的
	//pool.Close()
	connPool2 := pool.Get()

	_, err = connPool2.Do("Set", "namePool2", "xining")
	if err != nil {
		fmt.Println("connPool2.Do err= ", err)
		return
	}
	//取出
	rp2, err := redis.String(connPool2.Do("Get", "namePool2"))
	if err != nil {
		fmt.Println("connPool2.Do err= ", err)
		return
	}
	fmt.Println("rp2= ", rp2)
	fmt.Println("connPool2", connPool2)

	//因为返回r是interface{}
	//因为name对应的值是string，因此我们需要转换
	//nameString := r.(string)

}
