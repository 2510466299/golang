package main

import "fmt"

type person struct {
	name string
	age  int32
}

func person_generate(name string, age int32) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := person_generate("undersdong", 21)
	fmt.Printf("%v,%#v\n", p1, p1)
}
