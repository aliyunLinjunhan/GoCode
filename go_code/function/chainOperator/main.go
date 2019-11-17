package main

import (
	"fmt"
	"strings"
)

// 链式处理——操作与数据分离的设计技巧


// 自定义一个移除字符串前缀的处理函数
func removePrefix(str string) string {

	return strings.TrimPrefix(str, "go")
}

// 字符串处理函数， 传入字符串切片和处理链
func StringProccess(list []string, chain []func(string) string){

	//遍历每一个字符
	for index, str := range list {

		result := str

		for _, proc := range chain {

			result = proc(result)
		}

		list[index] = result
	}

}

func main() {

	// 未进行处理的字符串
	list := []string{
		"go scanner",
        "go parser",
        "go compiler",
        "go printer",
        "go formater",
	}

	// 处理函数链
	chain := []func(string) string {
		removePrefix,
        strings.TrimSpace,
        strings.ToUpper,
	}

	StringProccess(list, chain)
	fmt.Println(list)
	
}