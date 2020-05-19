package main

import "fmt"

func main(){
	fmt.Printf("\n")

	var s []int
	printSlice2(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice2(s)

	// the slice grows as needed.
	s = append(s, 1)
	printSlice2(s)

	// we can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice2(s)
}

func printSlice2(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
