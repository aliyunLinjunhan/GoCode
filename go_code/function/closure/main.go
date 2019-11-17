package main

import "fmt"

func Accumulate(value int) func() int {

	return func() int{
		
		value++

		return value
	}
}

func playerGen(name string) func() (string, int){

	hp := 150

	return func() (string, int){
		
		return name, hp
	}
}

func main() {
	
	str := "hello world"

	foo := func() {

		str = "hello dude"
	}

	foo()
	fmt.Println(str)

	accumulate := Accumulate(1)

	fmt.Println(accumulate())
	fmt.Println(accumulate())

	fmt.Printf(" %p\n", accumulate)

	accumulate2 := Accumulate(1)
	fmt.Println(accumulate2())

	fmt.Printf(" %p\n", accumulate2)
	// 对比输出的日志发现 accumulator 与 accumulator2 输出的函数地址不同

	// 创建一个游戏玩家
	generator := playerGen("linjunhan")

	fmt.Println(generator())

}