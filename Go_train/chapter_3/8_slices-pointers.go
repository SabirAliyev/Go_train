package main

import "fmt"

func main(){
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Andy",
	}
	fmt.Println("\n")
	fmt.Println(names)

	a := names[0:2]	// slice
	b := names[1:3] // slice
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}