package main

import (
	"fmt"
	"time"
	"sync"
	
)

var (
	myMap = make(map[int]int, 10)
	lock sync.Mutex  // 声明一个互斥🔓
)

func test(n int) {

	res := 1
	for i:=1;i<=n;i++ {
		res = res * i
	}
	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

func main() {

	for i:=1;i <= 200;i++ {

		go test(i)
	}

	time.Sleep(time.Second * 15)
	fmt.Println(myMap)
}