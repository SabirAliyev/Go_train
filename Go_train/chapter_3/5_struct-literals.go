package main

import "fmt"

type Omega struct {
	X, Y int
}

var (
	v1 = Omega{1,2}
	v2 = Omega{X: 1}
	v3 = Omega{}
	p = &Omega{1,2}
)

func main(){
	fmt.Println("\n")
	fmt.Println(v1, v2, v3, p)
}
