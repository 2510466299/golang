package main

import "fmt"

func main() {
	// byte=uint8   ASCIll
	// rune=int32
	s := "this is 中文示例"
	fmt.Println(s)
	for i, r := range s {

		fmt.Printf("%d,%c\n", i, r)
	}
}
