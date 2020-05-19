package main

import "fmt"

type Vertex4 struct {
	Lat, Long float64
}

var m map[string]Vertex4

func main(){
	fmt.Println("\n")
	m = make(map[string]Vertex4)
	m["Bell Labs"] = Vertex4{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
