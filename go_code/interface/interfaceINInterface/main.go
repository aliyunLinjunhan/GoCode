package main

import _ "fmt"

type A interface {
	A01()
}

type B interface {
	A02()
}

type C interface {
	A
	B
	A03()
}

type Stu struct {

}

func (stu Stu) A01(){

}
func (stu Stu) A02(){

}
func (stu Stu) A03(){

}

func main() {
	
	s := new(Stu)
	s.A03()
}
