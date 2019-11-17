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

	// 3.读取数据string类
	result, err := redis.String(conn.Do("Get", "name")) 
	//返回的result是一个空接口，需要转换为字符串，使用redis自带的字符转换方法

	if err != nil {
		fmt.Println("get err=", err)
		return 
	}
	fmt.Println("result=", result)

	// 4.操作Hash类数据
	// 存入数据
	_ , err = conn.Do("HSet", "user1", "name", "linjunhan林军韩", "age", "20")
	if err != nil {
		fmt.Println("HSet linjunhan err", err)
	}

	// 读取数据
	result, err = redis.String(conn.Do("HGet", "user1", "age"))
	if err != nil {
		fmt.Println("HGet err=", err)
		return 
	}
	fmt.Println("result=", result)

	// 5.批量set、get操作数据
	// 批量存入数据
	_ , err = conn.Do("HMSet", "user2", "name", "hhhhhhhh", "age", "25")
	if err != nil {
		fmt.Println("HMSet linjunhan err", err)
	}

	// 批量读取数据
	result2, err := redis.Strings(conn.Do("HMGet", "user2", "name", "age")) 
	//注意这里要使用redis.Strings, 这个结果和上面结果不一样
	//返回的是一个切片
	if err != nil {
		fmt.Println("HMGet err=", err)
		return 
	}
	for i, v := range result2 {
		fmt.Println("i ,v=", i, v)
	}




}