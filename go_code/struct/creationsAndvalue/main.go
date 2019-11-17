package main

import "fmt"

type Point struct {

	x int
	y int
}

type People struct {

	name string
	child *People
}

func main() {

	// 使用键值对初始化结构体
	ins := &Point{
		x : 10,
		y : 10,
	}
	fmt.Println(ins)

	relation := People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
				child: nil,
			},
		},
	}

	fmt.Println(relation)

	// 使用多值列表初始化结构体
	p := Point{10, 10}
	fmt.Println(p)

	// 初始化匿名结构体，无需通过type
	i := &struct {
		name string
		height int
		weight int
	}{
		name: "林军韩",
		height: 179,
		weight: 70,
	}
	fmt.Println(i)
}