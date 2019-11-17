package main
import (
	"fmt"
	"time"
)

func main() {

	var arr []int
	var flag bool
	start := time.Now().Unix()
	for num:=1; num<=800000; num++ {
		
		flag = true
		for i:=2 ; i < num; i++ {
			if num % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			arr = append(arr, num)
			fmt.Printf("素数是%v \n", num)
		}
	}
	end := time.Now().Unix()
	fmt.Println("所耗费时间为:", end - start, "s")
}