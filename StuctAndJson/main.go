package main

import "fmt"

type Student struct {
	Name string
	No   int
}

type Class struct {
	Title   string
	Studets []Student
}

func New_Student(name string, no int) Student {
	return Student{
		Name: name,
		No:   no,
	}
}
func main() {
	class := &Class{
		Title:   "good class",
		Studets: make([]Student, 0, 10),
	}
	for i := 0; i < 10; i++ {
		tmp := New_Student(fmt.Sprintf("stu%02v\n", i), i)
		class.Studets = append(class.Studets, tmp)
	}
	fmt.Printf("class:%#v\n", class)
	//json序列化
}
