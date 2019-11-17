package main

import (
	"fmt"
)

func main() {
	// 数组初始化赋值
	var a [3]int = [3]int{0, 1, 2}
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	// 对指定位置进行赋值
	var b [3]int = [3]int{0:1, 2:4}

	// 使用...来进行数组的初始化
	q := [...]int{0, 1, 2, 3, 5}
	fmt.Println(q[3])
	fmt.Println(b)

	for k, v := range q {
		fmt.Println(k, v)
	}

	// 声明多维数组
	var c [3][4]int = [3][4]int{{0, 1, 3}, {31, 32, 42}}
	fmt.Println(c)
	d := [...][4]int{1:{0, 3, 4, 4}, 2:{32, 3, 2, 4}}
	e := [...][4]int{{3, 32, 4, 4}}
	fmt.Println(d)
	fmt.Println(e)
}