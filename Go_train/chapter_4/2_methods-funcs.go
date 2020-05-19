package main

import (
	"fmt"
	"math"
)

func Square(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	g := Vertex{3, 4}
	fmt.Println("\n", Square(g))
}
