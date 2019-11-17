package main
import (
	"runtime"
	"fmt"
)

func main() {

	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	//可以设置使用多个cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}