package main

// 通过golang对redis操作，还可以通过redis连接池
// 1. 事先初始化一定数量的连接，放入到连接池
// 2. 当Go需要操作redis时，直接从redis取出链接即可
// 3. 这样可以节省临时获取redis链接的时间

// 使用技巧
// 1. 初始化连接池
// 2. 核心代码
// var pool *redis.Pool
// pool = &redis.Pool{
// 	Maxldle: 8,	// 最大空闲连接数
// 	MaxActive: 0,	// 表示和数据库的最大连接数，0表示没有限制
// 	ldleTimeout: 100,	// 最大空闲时间
// 	Dial: func() (redis.Conn, error){ // 初始化链接的代码
// 		return redis.Dial("tcp", "localhost:6379")
// 	}
// }

import(
	"github.com/garyburd/redigo/redis"
	"fmt"
)

// 定义一个全剧pool
var pool *redis.Pool

//当启动程序时，就初始化连接池
func init() {

	pool = &redis.Pool{
		MaxIdle: 8,	// 最大空闲连接数
		MaxActive: 0,	// 表示和数据库的最大连接数，0表示没有限制
		IdleTimeout: 100,	// 最大空闲时间
		Dial: func() (redis.Conn, error){ // 初始化链接的代码
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {

	// 从连接池中取出一个链接
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Set", "cat", "pangzei")
	if err != nil{
		fmt.Println("set err", err)
	}

	r, err := redis.String(conn.Do("Get", "cat"))
	if err != nil{
		fmt.Println("get err", err)
	}
	fmt.Println(r)
}