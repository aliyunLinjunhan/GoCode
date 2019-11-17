package main

// Go语言的匿名函数支持随用随调的原理
// 定义匿名函数的格式：
// func(参数列表)(返回参数列表){
//     函数体
// }
import "fmt"

func main() {

	// 在定义时调用匿名函数
	func(data int) {
		fmt.Println(data)
	}(100)

	// 将匿名函数赋值给变量
	f := func(data int) {
		fmt.Println(data)
	}

	f(100)
}
