package main

import "fmt"

type shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func PrintArea(s shape) {
	fmt.Println("Area : ", s.Area())
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 10, Height: 20}

	PrintArea(c)
	PrintArea(r)
}
