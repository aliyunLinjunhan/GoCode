package main

import (
	"fmt"
)

type Cat struct {
	Name string 
	AGe int
}

func main() {

	// 定义一个存放任意类型的管道 3个数据
	allChan := make(chan interface{}, 3)

	allChan<- 10
	allChan<- "tom jack"
	cat := Cat{"小花猫", 4}
	allChan<- cat

	// 我们希望获得到管道中的第三个元素
	<-allChan
	<-allChan

	newCat := <-allChan

	fmt.Printf("newCat=%T, newCat=%v\n", newCat, newCat)
	a := newCat.(Cat)  // 使用类型断言
	fmt.Printf("newCat=%v\n", a.Name)
}