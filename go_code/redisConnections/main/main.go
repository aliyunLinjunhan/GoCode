package main

import(
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {
	
	// 通过go向redis写入数据和读取数据
	// 1. 连接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis conn err", err)
		return 
	}
	defer conn.Close()
	fmt.Println("conn succ.......", conn)
	// 2.通过Go向redis写入数据 string [key-value]
	_ , err = conn.Do("Set", "name", "linjunhan林军韩")
	if err != nil {
		fmt.Println("Set linjunhan err", err)
	}

	// 3.读取数据
	result, err := redis.String(conn.Do("Get", "name")) 
	//返回的result是一个空接口，需要转换为字符串，使用redis自带的字符转换方法
	
	if err != nil {
		fmt.Println("set err=", err)
		return 
	}
	fmt.Println("result=", result)

}