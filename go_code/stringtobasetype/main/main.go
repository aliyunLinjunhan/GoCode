package main

import(
	"fmt"
	"strconv"
)

func main() {
	var str string = "true"
	var n1 bool
	n1, _ = strconv.ParseBool(str)
	fmt.Printf("n1 type : %T, n1: %t\n", n1, n1)

	var shu string = "321321"
	var n2 int64
	n2, _ = strconv.ParseInt(shu, 10, 64)
	fmt.Printf("n2 type : %T, n2: %d\n", n2, n2)
	// 在无法正常转换的情况下会将它转为0
	n2, _ = strconv.ParseInt(str, 10, 64)
	fmt.Printf("n2 type : %T, n2: %d\n", n2, n2)

	var flshu string = "32.4412"
	var n3 float64
	n3, _ = strconv.ParseFloat(flshu, 10)
	fmt.Printf("n3 type : %T, n3: %f\n", n3, n3)

}