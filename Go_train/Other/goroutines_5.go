package main

import (
	"fmt"
	"time"
)

// for read only function:
func fout(out chan<-int){
	out <- 5
}

func fin(in <-chan int){
	fmt.Printf("%d\n", <-in)
}

func main(){
	var ch = make(chan int)
	go fin(ch)
	go fout(ch)

	time.Sleep(5 * time.Second)
}
