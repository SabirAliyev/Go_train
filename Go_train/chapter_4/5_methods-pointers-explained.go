package main

import (
	"fmt"
	"math"
)

type Vertex4 struct {
	X, Y float64
}

func aBs(v Vertex4) float64 {
	return  math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex4, f float64) { // remove * to see another result.
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex4{3, 4}
	Scale(&v, 10)
	fmt.Println("\n", aBs(v))
}
