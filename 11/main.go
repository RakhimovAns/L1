package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3}

	set := make(map[int]struct{})
	for _, v := range a {
		set[v] = struct{}{}
	}

	var answer []int
	for _, v := range b {
		if _, exists := set[v]; exists {
			answer = append(answer, v)
		}
	}

	fmt.Println(answer)
}
