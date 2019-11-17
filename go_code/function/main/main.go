package main

import "fmt"

// 以下是函数参数传递效果
type Data struct {

	complax []int

	instance InnerData

	ptr *InnerData
}

type InnerData struct {
	a int
}

func passByValue(inFunc Data) Data {

	fmt.Printf("inFunc value: %+v\n", inFunc)
	fmt.Printf("inFunc ptr: %p\n", &inFunc)

	return inFunc
}

func main() {

	// 准备输入结构的函数体
	in := Data{
		complax: []int{1, 2, 3},
		instance: InnerData{5},
		ptr: &InnerData{1},
	}

	//输入结构的成员情况和指针
	fmt.Printf("输入结构的内容: %+v, 指针地址：%p \n", in, &in)

	out := passByValue(in)

	//输出结构的成员情况和指针
	fmt.Printf("输入结构的内容: %+v, 指针地址：%p \n", out, &out)

}