package main

import (
	"fmt"
)

func logger(prefix string) func(s string){
	return func(s string){
		fmt.Printf("[%s] %\n", prefix, s)
	}
}

func main(){
	warn := logger("WARN")
	err := logger("ERROR")
	warn("TEST")
	err("Tiёsto")
	err("Markus Schulz")
}

