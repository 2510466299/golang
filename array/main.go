package main

import "fmt"

func main() {
	var arr = [...]int{1, 2, 3, 45, 6}

	fmt.Println(arr)
	/*
		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i])
		}
	*/
	for index, value := range arr {
		fmt.Println(index, value)
	}
}
