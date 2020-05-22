package main

import (
	"fmt"
	"io"
)

type Exchanger interface {
	Exchange()
}

type StringPair struct {
	first, second string
}

func (pair *StringPair) Exchange(){
	pair.first, pair.second = pair.second, pair.first
}


func (pair *StringPair) Read(data []byte) (n int, err error) {
	if pair.first == "" && pair.second == "" {
		return 0, io.EOF
	}
	if pair.first != "" {
		n = copy(data, pair.first)
		pair.first = pair.first[n:]
	}
	if n < len(data) && pair.second != "" {
		m := copy(data[n:], pair.second)
		pair.second = pair.second[m:]
		n += m
	}
	return n, nil
}

type Point [2] int
func (point *Point) Exchange() { point[0], point[1] = point[1], point[0] }

func F () {
	jekyll := StringPair{"Henry", "Jekyll"}
	hyde := StringPair{"Edward", "Hyde"}
	point := Point{5, -3}
	fmt.Println("Before:  ", jekyll, hyde, point)
	jekyll.Exchange() 	// Интерпретируется как: (&jekyll).Exchange()
	hyde.Exchange() 	// Интерпретируется как: (&hyde).Exchange()
	point.Exchange() 	// Интерпретируется как: (&point).Exchange()
	fmt.Println("After #1:", jekyll, hyde, point)
	exchangeThese(&jekyll, &hyde, &point)
	fmt.Println("After #2:", jekyll, hyde, point)
	// Before: "Henry"+"Jekyll" "Edward"+"Hyde" [5 -3]
	// After #1: "Jekyll"+"Henry" "Hyde"+"Edward" [-3 5]
	// After #2: "Henry"+"Jekyll" "Edward"+"Hyde" [5 -3]
}

func exchangeThese(exchangers ...Exchanger) {
	for _, exchanger := range exchangers {
		exchanger.Exchange()
	}
}

func main(){
	fmt.Println("\n")
	F()
}