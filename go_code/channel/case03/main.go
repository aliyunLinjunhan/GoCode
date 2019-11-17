package main

import(
	"fmt"
)

func main() {

	myChan := make(chan int, 3)
	myChan<- 100
	myChan<- 200
	close(myChan)  // 关闭管道,关闭后，就不能向channel写数据，但是仍然可以从该channel读取数据。

	fmt.Println("mychan:", myChan)

	// channel支持for-range遍历
	// 如果channel没有关闭，则会出现deadlock死锁
	// 如果channel有关闭，则能全部整车昂遍历
	
	// 遍历向管道输入数据
	intchan2 := make(chan int, 100)
	for i:=0;i<100;i++ {
		intchan2<- i*2
	}

	close(intchan2)
	for V := range intchan2 {
		fmt.Printf("v=%v\t", V)
	}
	fmt.Println()

}