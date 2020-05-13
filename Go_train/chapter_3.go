package main

import "fmt"

func poinerUsage(){
	i, j := 42, 2701

	p := &i			// point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21 		// set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j 			// point to j
	*p = *p / 37 	// divide j through the pointer
	fmt.Println(j)  // see the new value of j
}

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1,2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p = &Vertex{1,2}
)

func Array(){
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func Slice(){
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

func main(){
	poinerUsage()

	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	p := &v
	p.X = 1e9 // (*p).X
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)

	Array()

	Slice()
}
