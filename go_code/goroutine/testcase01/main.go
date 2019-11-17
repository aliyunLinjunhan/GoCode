package main
import(
	"fmt"
	"time"
)

func test() {

	for i := 0;i < 10;i++ {

		fmt.Println("hello world!!")
		time.Sleep(time.Second)
	}
}

func main() {
	
	go test() //开启了一个携程
	for i := 0; i < 10; i++ {
		fmt.Println("hello go!!")
		time.Sleep(time.Second)
	}
}