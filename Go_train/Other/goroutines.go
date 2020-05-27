package main

import "fmt"

func main(){
	// go 	fmt.Printf("Goroutines: %d", runtime.NumGoroutine())

	var ch = make(chan struct{})

	go func(){
		fmt.Printf("Hello")
		ch <- struct{}{}
	}()
	<- ch

}

