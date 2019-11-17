package main

import "fmt"

type Cat struct {
	Color string
	Name string
}

type BlackCat struct {
	Cat
}

func NewCatByName(name string) *Cat{
	
	return &Cat{
		Name: name,
	}
}

func NewCatByCOlor(color string) *Cat{

	return &Cat{
		Color: color,
	}
}
// Go语言方法和接收器
type Bag struct {
	items []int
}

// 将一个物品放入背包的过错, 下面这是面向过程语言的写法
func Insert(b *Bag, itemid int) {

	b.items = append(b.items, itemid)
}

// Go语言构造方法的实现
func (b *Bag) Insert(itemid int){

	// 上方的(b *Bag)表示接收器，可以接受指针也可以接收实例
	b.items = append(b.items, itemid)
}

// 指针类型的接收器和非指针类型的接收器是有区别的 
func main() {
	blackCat := NewCatByCOlor("black")
	fmt.Println(blackCat)

	bag := new(Bag)
	Insert(bag, 1001)
	println(bag)

	b := *new(Bag)
	b.Insert(1001)
	fmt.Println(b)
}