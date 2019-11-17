package main

import (
	"fmt"
	_ "time"
	"sync"
)

var myMap = make(map[int]int, 2000)
var lock sync.Mutex

func writeChannel(intChan chan int){

	for i:=1;i<=2000;i++{
		intChan<- i
	}
}

func readChannel(intChan chan int, exitChan chan bool) {

	for {
		k, _ := <-intChan
		lock.Lock()
		myMap[k] = sum(k)
		lock.Unlock()
		if k == 2000 {
			// 读完数据后，任务完成，
			exitChan<- true
			close(exitChan)
		}
	}
}

func sum(n int) int{

	res := 0
	for i:=1;i<=n;i++ {
		res = res + i
	}
	return res
}

func main(){

	numChan := make(chan int, 2000)
	exitChan := make(chan bool, 1)
	// resChan := make(chan map[int][int], 20000)
	go writeChannel(numChan)
	go readChannel(numChan,exitChan)
	go readChannel(numChan,exitChan)
	go readChannel(numChan,exitChan)
	go readChannel(numChan,exitChan)
	go readChannel(numChan,exitChan)
	go readChannel(numChan,exitChan)
	go readChannel(numChan,exitChan)
	go readChannel(numChan,exitChan)

	// time.Sleep(time.Second * 5)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
	lock.Lock()
	fmt.Println(myMap)
	lock.Unlock()
	

}