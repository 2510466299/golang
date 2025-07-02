package main

import "fmt"

func change(s ...string) {
	s[0] = "go"
	s = append(s, "playground!")
	s = append(s, "give me passion!")
	fmt.Println(s)

}
func main() {
	s := []string{"hello", "world"}
	change(s...)
	fmt.Println(s)
}
