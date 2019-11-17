package main

import (
	"fmt"
)

func main() {

	// 1.演示管道的示例
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)

	// 2.向管道中写入数据
	intChan<- 10
	num := 211
	intChan<- num

	// 3.查看管道的长度和容量，不能超过容量
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))

	// 4.从管道中读取数据
	var num2 int
	num2 = <-intChan
	<-intchan 	//直接扔掉管道中的东西 
	fmt.Println("num2 = ", num2)
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))

	// 5.在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取会报错deadlock

	// 6.
}