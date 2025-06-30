package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// map
func main() {
	a := make(map[string]int, 10)
	a["数学"] = 100
	a["语文"] = 90
	a["英语"] = 80
	fmt.Println(a)
	v, ok := a["语文"]
	fmt.Println(v, ok)
	v, ok = a["物理"]
	fmt.Println(v, ok)
	delete(a, "数学")

	mapScore := make(map[string]int, 50)
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		mapScore[key] = value
	}
	sliceScore := make([]string, 0, 100)
	for key := range mapScore {
		sliceScore = append(sliceScore, key)
	}
	sort.Strings(sliceScore)
	for _, key := range sliceScore {
		fmt.Println(key, mapScore[key])
	}

	//值为 map 的切片
	mapSlice := make([]map[string]int, 8, 8)
	for index, _ := range mapSlice {
		mapSlice[index] = make(map[string]int, 8)
		mapSlice[index]["张三"] = rand.Intn(100) // 每个 map 的值都为随机数
	}
	fmt.Println(mapSlice[0])
	//map'value is slice
	sliceMap := make(map[string][]int, 8)
	value, ok := sliceMap["golang"]
	if ok {
		fmt.Println(value)
	} else {
		sliceMap["golang"] = make([]int, 8)
		sliceMap["golang"][0] = 100
		sliceMap["golang"][1] = 200
		sliceMap["golang"][2] = 300
		fmt.Println(sliceMap["golang"])
	}

	for k, v := range sliceMap {
		fmt.Println(k, v)
	}
}
