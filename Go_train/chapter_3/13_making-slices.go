package main

import "fmt"

func main(){
	a := make([]int, 5)
	fmt.Println("\n")
	printMadeSlices("a", a)

	b := make([]int, 0, 5)
	printMadeSlices("b", b)

	c := b[:2]
	printMadeSlices("c", c)

	d := c[2:5]
	printMadeSlices("d", d)
}

func printMadeSlices(s string, x []int){
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
