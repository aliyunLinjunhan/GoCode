package main

import(
	"fmt"
	"unsafe"
	"strconv"
)
	

func main() {
	var n1 int32 = 10
	var n2 float32 = 32.32
	var n3 bool = true

	var str string = fmt.Sprintf("%d", n1)
	fmt.Printf("n1:%q, type of %T:\n", str, str)
	str = fmt.Sprintf("%f", n2)
	fmt.Printf("n2:%q, type of %T:\n", str, str)
	str = fmt.Sprintf("%t", n3)
	fmt.Printf("n3:%q, type of %T:\n", str, str)

	var n4 int = 99
	var n5 float32 = 32.43
	fmt.Printf("n4:%T, n5:%T, n4:%d\n", n4, n5, unsafe.Sizeof(n4))

	str = strconv.FormatInt(int64(n4), 10)
	fmt.Printf("n4:%q, type of %T:\n", str, str)
	// 参数： float64类型的数， 数据的输出格式， 精度， 源数据格式
	str = strconv.FormatFloat(float64(n5), 'f', 10, 32)
	fmt.Printf("n5:%q, type of %T:\n", str, str)
}