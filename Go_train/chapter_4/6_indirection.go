package main

import "fmt"

type Vertux struct {
	X, Y float64
}

func (v *Vertux) Scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertux, f float64)  {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main(){
	v := Vertux{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertux{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println("\n", v, p)
}
