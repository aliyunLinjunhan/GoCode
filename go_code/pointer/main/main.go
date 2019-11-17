package main

import (
	"fmt"
	"../model"
)

func main() {
	// 声明一个变量
	var i int = 5
	fmt.Printf("i的内存地址%v\n", &i)

	// 声明一个指针
	var ptr *int = &i
	fmt.Printf("ptr的内容：%v\n", ptr)
	fmt.Printf("ptr的地址: %v\n", &ptr)
	fmt.Printf("ptr所执行的空间内容: %v\n", *ptr)


	var num int = 5
	var ptrNum *int = &num
	fmt.Printf("num原来的值为%v\n", num)
	*ptrNum = 6
	fmt.Printf("修改后num的值为%v\n", num)

	fmt.Println(model.StuNum)
}