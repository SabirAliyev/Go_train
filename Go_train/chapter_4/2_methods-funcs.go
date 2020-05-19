package main

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	X, Y float64
}

func Square(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	g := Vertex2{3, 4}
	fmt.Println("\n", Square(g))
}
