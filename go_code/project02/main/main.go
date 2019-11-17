package main
import "fmt"

const pi = 3.1415926

const (
	statusok = 200
	notFound = 404
)

// 批量声明，没有赋值默认和上面一样
const (
	n1 = 200
	n2
	n3
)

// iota
const (
	a1 = iota
	a2 = iota
	_ = iota
	a3 = iota
)

// 定义数量级
const (
	_ = iota
	KB = 1 << (10 * iota)   // 1左移10位 二进制： 10000000000 十进制： 1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main(){
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	// 十进制
	var a int = 10
	fmt.Printf("%d\n", a)
	fmt.Printf("%b\n", a) //二进制输出

	// 八进制 以0开头
	var b int = 077
	fmt.Printf("%o\n", b)
	fmt.Printf("%d\n", a)

	// 十六进制
	var c int = 0xff
	fmt.Printf("%x\n", c)
	fmt.Printf("%T\n", c)  // 输出数据类型

	// 浮点数
	d := 1.3
	fmt.Println(d)

	if a >10 {
		
	} else if {

	}
}