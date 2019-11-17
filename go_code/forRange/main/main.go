package main

import "fmt"

func main() {
	silce := []int{0, 1, 2, 3}
	for index, value := range silce {
		fmt.Printf("value: %d Valuer-addr: %X, Elemarr: %X\n", value, &value, &silce[index])
	}

	silce2 := [][]int{{10}, {100, 200}}
	fmt.Println(silce2)
	silce2[0] = append(silce2[0], 20)
	fmt.Println(silce2)
}