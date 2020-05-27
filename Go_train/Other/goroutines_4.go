package main

import "fmt"

func main(){
	var ch = make(chan int)

	go func(){
		fmt.Printf("Hello")
		//ch <- 1
		close(ch)
	}()

	x, ok := <- ch
	fmt.Printf("  x: %v, ok: %v", x, ok)
}
