package main

import "fmt"

type Animal struct {
	name  string
	age   int
	voice string
}

func (a *Animal) Voice() {
	fmt.Printf("%v会%v%v%v\n", a.name, a.voice, a.voice, a.voice)
}

type Dog struct {
	feet int
	*Animal
}

func (d *Dog) Feets() {
	fmt.Printf("%v有%v条腿\n", d.name, d.feet)
}

func main() {
	d := &Dog{
		feet: 4,
		Animal: &Animal{
			name:  "乐乐",
			age:   2,
			voice: "汪",
		},
	}
	d.Voice()
	d.Feets()
}
