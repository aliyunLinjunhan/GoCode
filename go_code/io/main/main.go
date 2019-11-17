package main

import (
	"fmt"
	_ "bufio"
	_ "os"
	"io/ioutil"
)

func main() {

	// 带缓冲区的文件读取方式
	// file, err := os.Open("a.txt")
	// if err != nil {
	// 	fmt.Println("open file err=", err)
	// }
	// defer file.Close()

	// reader := bufio.NewReader(file)
	// for {
	// 	str, err := reader.ReadString('\n')
	// 	if err == io.EOF{
	// 		break
	// 	}
	// 	fmt.Print(str)	
	// }

	// 使用ioutil一次将文件读到内存中，这种适用于文件不太大的情况
	file := "a.txt"
	f, err := ioutil.ReadFile(file)
	if err != nil{
		break
	}
	fmt.Print(str)	

}