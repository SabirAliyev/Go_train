package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) String() string{
	return fmt.Sprintf("\n %v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Artur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

