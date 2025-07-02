package main

import "fmt"

type Mytype int64
type Person struct {
	name string
	age  int64
}

func main() {
	//global struct
	var i Mytype = 10
	fmt.Printf("%T,%v\n", i, i)
	var p Person
	p.name = "understood"
	p.age = 21
	fmt.Printf("%T,%#v\n", p, p)

	//tmp struct
	var tmp struct {
		x, y int32
	}
	tmp.x = 1
	tmp.y = 2
	fmt.Println(tmp)
	//new struct
	p1 := new(Person)
	p1.age = 22
	p1.name = "undersdong"
	fmt.Printf("%T,%#v\n", p1, p1)

	//ptr struct
	p2 := &Person{}
	p2.age = 21
	p2.name = "undersdong"
	fmt.Printf("%T,%#v\n", p2, p2)

	//ptr struct
	p3 := &Person{
		"undersdong",
		21,
	}
	fmt.Printf("%T,%#v\n", p3, p3)
}
