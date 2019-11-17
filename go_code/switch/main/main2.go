package main

import (
	"fmt"
)

func main() {
	// var c byte
	// fmt.Println("请输入小写字母：")
	// fmt.Scanf("%c", &c)

	// switch c {
	// 	case 'a':
	// 		fmt.Println("A")
	// 	case 'b':
	// 		fmt.Println("B")
	// 	case 'c':
	// 		fmt.Println("C")
	// 	case 'd':
	// 		fmt.Println("D")
	// 	case 'e':
	// 		fmt.Println("E")
	// 	default:
	// 		fmt.Println("other")
	// }

	var score float32
	for {
		fmt.Println("请输入您的成绩")
		fmt.Scanf("%f", &score)
		if score >= 0 && score <= 100 {
			goto switch_begin
		} else {
			fmt.Println("您输入的成绩有误，请重新输入！！")
		}
	}

switch_begin:
	switch int(score / 60) {
		case 1:
			fmt.Println("合格！！")
		case 0:
			fmt.Println("不合格！！")
		default: 
			fmt.Println(score/60)
	}

}
