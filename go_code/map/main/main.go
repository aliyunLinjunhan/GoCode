package main

import "fmt"

func main() {
	var mapList map[string]int
	mapList = map[string]int{"wo": 100}
	mapList["432423"] = 312
	fmt.Println(mapList)
	mapSsign := make(map[string]float32, 10)
	mapSsign["dadas"] = 100
	fmt.Println(mapSsign)

	// 用切片作为map的值

	var silce []int
	silce = append(silce, 1, 2)
	mapSilce := map[int][]int{1 :silce}
	fmt.Println(silce)
	fmt.Println(mapSilce)
	
	// 使用for-range循环进行迭代
	for key, value := range mapList {
		fmt.Println(key, value)
	}

	// map的删除和清空
	// 使用 delete() 内建函数从 map 中删除一组键值对，delete() 函数的格式如下：delete(map, 键)
	scene := make(map[int]float32)
	scene[3] = 23.3
	scene[2] = 43.2
	scene[4] = 43.3
	fmt.Println(scene)
	delete(scene, 3)
	fmt.Println(scene)

	// 特别注意有sync.map，解决线程冲突问题
}