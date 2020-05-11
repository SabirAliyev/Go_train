package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
)

func add(x, y int) int {
	return x + y;
}

func swap(x, y string) (string, string) {
	return x, y
}

func split(sum int) (x, y int) {
	x = sum * 4/9
	y = sum - x
	return
}

var c, python, java bool
var k, j = 1, 2 //cannot use :=

var (
	ToBe	bool	= false
	MaxInt	uint64	= 1 << 64 - 1
	z		complex128 = cmplx.Sqrt(-5 + 12i)
)


func main() {
	fmt.Println("\nThe Number is:", rand.Intn(10))

	fmt.Println(add (42, 13))

	a, b := swap("Connection", "detected")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)

	c, python, java := true, false, "no!"
	fmt.Println(k, j, c, python, java)

	fmt.Printf("\nType: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

}

