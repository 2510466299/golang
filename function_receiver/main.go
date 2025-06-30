package main

import "fmt"

type Person struct {
	name string
	age  int32
}

func New_person(name string, age int32) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// define function_receiver

func (p Person) Dream() {
	fmt.Printf("%v,is learning golang\n", p.name)
}
func (p Person) Showage() {
	fmt.Printf("%v'age is %v\n", p.name, p.age)
}

func (p *Person) SetAge(newAge int32) {
	fmt.Printf("%v'age is %v\n", p.name, p.age)
	p.age = newAge
	fmt.Printf("%v'age is %v\n", p.name, p.age)
}
func main() {
	p := &Person{
		"undersdong",
		21,
	}
	p.Dream()
	p.Showage()
	p.SetAge(23)
}
