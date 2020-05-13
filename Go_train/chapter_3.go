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

func StringSlice(){
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Andy",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func SliceLiterals(){
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false}
	fmt.Println(r)

	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func SliceBounds(){
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func SliceLenCap(){
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	//Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func NilSlice(){
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("Slice is nil!")
	}
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

	fmt.Println("\nSlice:")
	Slice()

	fmt.Println("\nSlice with strings:")
	StringSlice()

	fmt.Println("\nSlice Literals:")
	SliceLiterals()

	fmt.Println("\nSlice Bounds:")
	SliceBounds()

	fmt.Println("\nSlice len, cap:")
	SliceLenCap()

	fmt.Println("\nNil Slice")
	NilSlice()
}
