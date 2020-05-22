package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main(){
	var a Abser
	f := MyFloat2(-math.Sqrt2)
	v := Vertox{3 ,4}

	a = f	// a MyFloat2 implements Abser
	a = &v  // a *Vertox implements Abser

	// In the following line, v is a Vertox (not *Vertox)
	// and does NOT implement Abser
	a = v

	fmt.Println("\n", a.Abs())
}

type MyFloat2 float64

func (f MyFloat2) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertox struct {
	X, Y float64
}

// Don`t use * before Vertox.
func (v Vertox) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}