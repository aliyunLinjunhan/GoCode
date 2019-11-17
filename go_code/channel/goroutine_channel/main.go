package main

import (
	"fmt"
	_ "time"
)

func writeData(intChan chan int) {
	
	for i:=1; i<= 100;i++ {
		// 放入数据
		intChan<- i
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到的数据=%v\n",v)
	}
	// 读完数据后，任务完成，
	exitChan<- true
	close(exitChan)
}

func main() {

	// 创建俩个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)  // 主线程立马就没有了，也不会去运行协程

	// time.Sleep(time.Second * 10)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}


}