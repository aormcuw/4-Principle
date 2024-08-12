package main

import "fmt"

type Speaker interface {
	Speak()
}

type Cat struct{}
type Cow struct{}

func (c Cat) Speak() {
	fmt.Println("Meow")
}

func (c Cow) Speak() {
	fmt.Println("Mooo")
}

func MakeSound(s Speaker) {
	s.Speak()
}

func main() {
	var cat Cat
	var cow Cow

	cat.Speak()
	cow.Speak()
}
