package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "this is a test string"
	fmt.Println(s1)
	s2 := strings.Split(s1, " ")
	fmt.Printf("%T\n", s2)
	fmt.Println(s2)
	s3 := strings.LastIndex(s1, "i")
	fmt.Printf("%T\n", s3)
	s4 := []string{"what", "do", "you", "want"}
	fmt.Println(strings.Join(s4, "-"))
}
