package main

import "fmt"

func main() {
	// 类型断言
	var x interface{}
	var b2 float32 = 1.1
	x = b2  //空接口可以接收任意类型

	// 带检测的类型断言
	if y, ok := x.(float64); ok{
		fmt.Println("convert success")
		fmt.Printf("y 的类型是 %T 值是=%v", y, y)
	}else{
		fmt.Println("convert fail")
	}
	fmt.Println("继续执行....")
}