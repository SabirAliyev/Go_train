package main

import (
	"fmt"
)

func main(){
	var ch = make(chan int, 3)

	for i := 0; i < 3; i++ {
		ch <- i
		fmt.Println("Recorded! ")
	}

	for j := 0; j < 4; j++ {
		fmt.Printf("Reading: %v; ", <-ch)
	}
}
