package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// For circle.
func SumOfTen() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

//Initialisation and finish blocks are optional.
func ForOptions() int {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	return sum
}

//For is White in Go
func WhileIsFor() int {
	sum := 1
	for sum < 1000{
		sum += sum
	}
	return sum
}

//Operator "if"
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// Short "if" instruction
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// Variables in "if" statement are accessible in "else" statement also.
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can`t use v here, though
	return lim
}

//Switch Case
func SwitchCase(){
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbssd,
		// plan9, windows...
		fmt.Printf("%s.", os)

	}
}

func SwitchCase2(){
	fmt.Println("When`s Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func Defer(){
	defer fmt.Println("World")
	fmt.Println("Hello")
}

func DeferStack(){
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main(){
	fmt.Println(SumOfTen())
	fmt.Println(ForOptions())
	fmt.Println(WhileIsFor())
	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
		)
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
	SwitchCase()
	SwitchCase2()
	Defer()
	DeferStack()
}



