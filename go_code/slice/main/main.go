package main

import (
	"fmt"
	_ "unsafe"
)

func main() {
	// 从数组中生成切片
	var a [3]int = [3]int{0, 1, 2}
	fmt.Println(a, a[1:2])

	var highRiseBUilding [30]int 
	for i:=0;i<30;i++ {
		highRiseBUilding[i] = i+1
	}
	fmt.Println(highRiseBUilding)

	// 输出前10个
	fmt.Println(highRiseBUilding[:10])
	// 输出15-25个
	fmt.Println(highRiseBUilding[14:25])
	// 输出最后10个
	fmt.Println(highRiseBUilding[20:])

	// 表示原有的切片
	fmt.Println(highRiseBUilding[:])

	// 重置切片
	fmt.Println(highRiseBUilding[0:0])

	// 声明一个字符串切片
	var strlist []string

	// 声明一个整形切片
	var intlist []string

	// 声明一个空切片
	var emptyLIst = []int{}
	fmt.Println(strlist, intlist, emptyLIst)


	//使用make()函数构造切片, 第一个参数是切片类型， 第二个是分配的元素个数， 第三个是预分配空间
	ai := make([]int, 2, 10)
	var bi []int = make([]int, 2, 10)
	fmt.Println(ai)
	fmt.Println(bi)

	// 使用append为切片尾部添加元素
	var as []int
	as = append(as, 1, 2, 3)
	fmt.Println(as)
	// 看看在不断累加元素后切片的底层变化
	// 当空间不够时，append进行俩倍扩容，并且会导致地址改变
	var numbers []int
	for i:=0;i<10;i++ {
		numbers = append(numbers, i+1)
		fmt.Printf("len:%d, cap:%d, %p\n", len(numbers), cap(numbers), numbers)
	}
	// 使用append为切片头部添加元素
	var ca []int = []int{0, 1, 2}
	ca = append([]int{0, 1}, ca...)
	ka := []int{0, 1, 2}
	ka = append([]int{0, 1}, ka...)
	fmt.Println(ca, ka)
	
	// 切片的拷贝
	silce1 := []int{1, 2, 3, 4, 5}
	silce2 := []int{32, 32, 32}
	fmt.Printf("silce1的地址%p\n", silce1)
	copy(silce1, silce2)
	fmt.Println(silce1)
	fmt.Printf("接受拷贝后的地址为%p\n", silce1)

	// 切片的删除， 如果删除的是前面的元素一定会导致地址的变化
	silce3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%p, %p\n", silce3, silce3[:4])

	i := []int{0, 1, 2}
	p := [3]int{0, 1, 2}
	fmt.Printf("i:%T, p:%T\n", i, p)

	// 注意slice和数组在声明时的区别：声明数组时，方括号内写明了数组的长度或使用...自动计算长度，而声明slice时，方括号内没有任何字符。
}