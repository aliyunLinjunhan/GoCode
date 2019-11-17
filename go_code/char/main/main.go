package main

import (
	"fmt"
	"unsafe"
)

func main(){
	
	var c1 byte = 'a' // 注意用单引号
	var c2 byte = '0'

	//当我们直接输出byte值，就是输出了对应的字符的码值
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	//如果想要输出字符，就必须使用格式化输出
	fmt.Printf("%c\n", c1)

	//当我要存储的字符超过了255， 可以采用int类型进行存储
	var c int = '北'
	fmt.Printf("%c\n", c)
	fmt.Printf("%d\n", c)

	var o = true
	fmt.Println("o:", o)
	fmt.Println("o:", unsafe.Sizeof(o))  //这个方法可以输出变量所占字节

	k := `
	package main

import (
	"fmt"
	"unsafe"
)

func main(){
	
	var c1 byte = 'a' // 注意用单引号
	var c2 byte = '0'

	//当我们直接输出byte值，就是输出了对应的字符的码值
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	//如果想要输出字符，就必须使用格式化输出
	fmt.Printf("%c\n", c1)

	//当我要存储的字符超过了255， 可以采用int类型进行存储
	var c int = '北'
	fmt.Printf("%c\n", c)
	fmt.Printf("%d\n", c)

	var o = true
	fmt.Println("o:", o)
	fmt.Println("o:", unsafe.Sizeof(o))  //这个方法可以输出变量所占字节
}`
	fmt.Println(k)
}