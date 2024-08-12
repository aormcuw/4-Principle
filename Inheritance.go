package main

import (
	"fmt"
	"strings"
)

type Order string
type Foot string

const (
	Carnivore Order = "Carnivore"
	Ominvore  Order = "Omnivore"
	Herbivore Order = "Herbivore"
)

const (
	Four Foot = "Four foots"
	Two  Foot = "Two foots"
)

type Animal struct {
	Name   string
	Orders []Order
	Foots  []Foot
}

// func (a *Animal) Speak() {
// 	fmt.Println("Animal sound")
// }

type Dog struct {
	Animal
	Breed string
}

type Cat struct {
	Animal
	Breed string
}

func (d *Dog) Speak() {
	fmt.Println("Woof")
}

func (c *Cat) Speak() {
	fmt.Println("Meow")
}

func NewAnimal(name string, orders []Order, foots []Foot, animalType string, breed string) interface{} {
	switch animalType {
	case "Dog":
		return Dog{
			Animal: Animal{
				Name:   name,
				Orders: orders,
				Foots:  foots,
			},
			Breed: breed,
		}
	case "Cat":
		return Cat{
			Animal: Animal{
				Name:   name,
				Orders: orders,
				Foots:  foots,
			},
			Breed: breed,
		}
	default:
		return nil
	}
}

func FormatSlice(slice interface{}) string {
	var sb strings.Builder
	switch v := slice.(type) {
	case []Order:
		for i, order := range v {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(string(order))
		}
	case []Foot:
		for i, foot := range v {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(string(foot))
		}
	}
	return sb.String()
}

func Describe(a interface{}) {
	switch v := a.(type) {
	case Dog:
		fmt.Printf("I am a Dog, i am a %s and i have %s\n", FormatSlice(v.Orders), FormatSlice(v.Foots))
	case Cat:
		fmt.Printf("I am a Cat, i am a %s and i have %s\n", FormatSlice(v.Orders), FormatSlice(v.Foots))
	default:
		fmt.Printf("i dunno \n")
	}
}

func main() {
	// d := Dog{
	// 	Animal: Animal{
	// 		Name:   "Buddy",
	// 		Orders: []Order{Carnivore},
	// 		Foots:  []Foot{Two}},
	// 	Breed: "Golden Retriever"}
	// c := Cat{
	// 	Animal: Animal{
	// 		Name:   "Pal",
	// 		Orders: []Order{Carnivore},
	// 		Foots:  []Foot{Four},
	// 	},
	// 	Breed: "Sphinx",
	// }
	d := NewAnimal("Buddy", []Order{Carnivore}, []Foot{Four}, "Dog", "Golden Retriever").(Dog)
	c := NewAnimal("Pal", []Order{Carnivore}, []Foot{Four}, "Cat", "Sphinx").(Cat)

	fmt.Println(d.Name)
	Describe(d)
	d.Speak()

	Describe(c)
	fmt.Println(c.Name)
	c.Speak()
}
