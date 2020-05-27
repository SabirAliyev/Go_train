package main

import (
	"fmt"
)

func main(){
	var ch = make(chan int, 3)

	for i := 0; i < 2; i++ {
		ch <- i
		fmt.Println("Recorded! ")
	}
	close(ch)

	for j := 0; j < 3; j++ {
		x, ok := <-ch
		fmt.Printf("x: %v, ok: %v\v", x, ok)
	}

	for j := 0; j < 4; j++ {
		fmt.Printf("len: %d, cap: %d", len(ch), cap(ch))
	}
}
