package main

import (
	"fmt"
	"container/list"
)

func main() {
	
	// 列表存在俩种声明方式New()和var
	// 列表与切片和 map 不同的是，列表并没有具体元素类型的限制，因此，列表的元素可以是任意类型.
	li := list.New()
	var li2 list.List

	// 向列表中插入元素
	li.PushBack("first")
	li.PushFront("123")
	li2.PushBack("second")
	fmt.Println(li.Front().Value, li2.Back().Value)

	// 便利列表
	for i:=li.Front();i != nil;i = i.Next() {
		fmt.Println(i.Value)
	}


}