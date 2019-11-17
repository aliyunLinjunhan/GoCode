package main

import "fmt"
type Point struct {

	x int
	y int
}

type player struct {
	name string
	HealthPoint int
	MagicPoint int
}

func main() {

	var p Point // 这个p是实例化对象
	p.x = 10
	p.y = 10
	fmt.Println(p)

	ins := new(Point) // 这个ins是一个指向结构体的指针
	ins.x = 9
	ins.y = 0
	fmt.Println(ins, *ins)

	player2 := *new(player) //这个声明方式出来的变量是结构体
	player2.name = "kuli"
	player2.HealthPoint = 10
	player2.MagicPoint = 10
	fmt.Println(player2)

	player3 := &player{} //这个声明取到一个指针
	player3.name = "sala"
	player3.HealthPoint = 10
	player3.MagicPoint = 10
	fmt.Println(player3)

	player1 := new(player)
	player1.name = "姚明"
	player1.HealthPoint = 10
	player1.MagicPoint = 10
	fmt.Println(player1)
}