package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func (p *Person) getAge() int {
	return p.age
}

func (p *Person) setAge(age int) {
	if age > 0 {
		p.age = age
	}
}

func main() {
	p := Person{Name: "Ais"}
	p.setAge(76)
	fmt.Println(p.Name)
	fmt.Println(p.getAge())
}
