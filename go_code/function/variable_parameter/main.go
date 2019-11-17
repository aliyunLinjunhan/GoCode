package main

import "fmt"

// 定义一个有可变参数的函数
func myfunc(args ...int) {

	// 上面这种参数的写法与 func myfunc(args []int) 一致， 但是如果采用这种写法，在调用时将不得不用切片进行传入
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func main() {
	myfunc(1, 2, 3)
	var a int = 54
	var b string = "hello"
	var c int64 = 43
	MyPrintf(a, b, c)
}

// 任意类型的可变参数
func MyPrintf(args ...interface{}) {
	
	for _, arg := range args {
		// 通过switch进行判断类型
		switch arg.(type) {
			case int:
				fmt.Println("这是一种int类型")
			case string:
				fmt.Println("这是一种string类型")
			case int64:
				fmt.Println("这是一种Int64类型")
			default:
				fmt.Println("无法识别类型")
		}
	}
}