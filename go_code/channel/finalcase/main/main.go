package main

import (
	"fmt"
	"time"
)

func main() {
	
	intChan := make(chan int, 10000)
	primeChan := make(chan int, 200000) //放入结果
	// 标识退出的管道
	exitChan := make(chan bool, 4) 

	// 开始人物时间
	start := time.Now().Unix()
	// 开启一个协程，向 intChan放入 1-8000个数
	go putNum(intChan)
	for i:=1;i<=8;i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	// 当我们从这个管道取出四个值，就可以放心跳出主线程
	go func(){
		for i := 0; i < 4 ;i++ {
			<-exitChan
		}
		close(primeChan)
		end := time.Now().Unix()
		fmt.Println("耗费时间为", end - start, "s")
	}()

	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
	}
	fmt.Println("主线程关闭")
}

// 向intChan 放入 1-8000个数
func putNum(intChan chan int) {

	for i:=1;i<=800000;i++ {
		intChan<- i
	}
	close(intChan)
}

// 从intChan中取出数据，并判断是否为素数，如果是就
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i:=2 ; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan<- num
			fmt.Printf("素数是%v \n", num)
		}
	}
	fmt.Println("有一个primeNum 携程因为取不到数据关闭了！！！")
	exitChan<- true
}