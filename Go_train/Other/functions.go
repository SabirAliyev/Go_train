package main

import "fmt"

func SquearesOfSumAndDiff(a int64, b int64) (s int64, d int64){
	x, y := a+b, a-b
	s = x * x
	d = y * y
	return // <=> returns s, d
}

func main() {
	a, b := SquearesOfSumAndDiff(10, 1)
	fmt.Println("\n", a, b)
}
