package main

import "fmt"

func main() {
	var name string
	var age byte
	var sal float32
	var isPsss bool


	// 先用scanln一行一行输入，向scanf中传递的是地址 要使用&进行秋地址运算
	// fmt.Println("名字：")
	// fmt.Scanln(&name)

	// fmt.Println("年龄：")
	// fmt.Scanln(&age)

	// fmt.Println("工资：")
	// fmt.Scanln(&sal)

	// fmt.Println("是否通过考试：")
	// fmt.Scanln(&isPsss)

	// fmt.Println(name, age, sal, isPsss)

	// 使用scanf进行输入，一次性全部放入变量
	fmt.Scanf("名字：%s, 年龄: %d, 工资: %f,是否通过考试: %t\n", &name, &age, &sal, &isPsss)
	fmt.Println(name, age, sal, isPsss)
}