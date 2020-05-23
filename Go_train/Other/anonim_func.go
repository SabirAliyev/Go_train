package main

import (
	"fmt"
	"sort"
)

func E (){
	people := []string{"Dave", "Bob", "Alice"}

	sort.Slice(people, func(i int, j int) bool {
		return len(people[i]) < len(people[j])
	})

	fmt.Println(people)
}

func main(){
	E()
}
