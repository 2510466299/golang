package main

import (
	"fmt"
	"math"
)

const (
	a = iota
	b = iota //如果常量的表达式没有写，默认和上一行一致
	c        //如果常量的表达式没有写，默认和上一行一致
	d
	_
	f
	g = 100
	h = iota
)

func main() {

	fmt.Println(a, " ", b, " ", c, " ", d, " ", f, " ", g, " ", h)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	s :=
		`
	this is a 
	double line 
	`
	fmt.Println(s)
	s1 := "this is s1"
	s2 := s1 + s
	fmt.Println(s2)

}
