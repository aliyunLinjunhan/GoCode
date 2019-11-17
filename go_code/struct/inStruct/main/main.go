package main

import "fmt"

// 车轮
type Wheel struct {
	Size int
}

// 引擎
type Engine struct {
	Power int
	Type string
}

type Car struct{
	Wheel
	Engine
}

func main() {

	c := Car{

		Wheel: Wheel{
			Size: 18,
		},

		Engine: Engine{

			Type: "1.4T",
			Power: 143,
		},
	}
	fmt.Printf("%+v\n", c)
}