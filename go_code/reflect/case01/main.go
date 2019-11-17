package main

import (
	"fmt"
	"reflect"
)

func resflectTest(a interface{}){

	res := reflect.TypeOf(a)
	rTp := reflect.ValueOf(a)
	v := rTp.Float()
	fmt.Println("res:", res, "rTp:", rTp)
	fmt.Println("v 的类型:", v)
}

func main(){

	var v float64 = 1.2
	resflectTest(v)

}