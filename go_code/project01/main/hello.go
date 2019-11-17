package main

import "fmt"

// 声明变量 单独声明
// var name string

// 批量声明， 默认值
var (
	name string //""
	age  int    // 0
	isok bool   // false
)

func main() {
	var test1 string
	fmt.Println("deew")
	test1 = "做个小饰演"
	fmt.Println(test1)

	// 声明变量同时赋值
	var age int = 11
	println(age)
	// 类型推导，自动识别变量类型
	var name = "hhhh"
	println(name)
	// 简短变量声明
	simplename := "ljh"
	println(simplename)
}
