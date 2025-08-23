package main

import (
	"fmt"
)

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})
	for _, w := range words {
		set[w] = struct{}{}
	}

	for w := range set {
		fmt.Print(w, " ")
	}
}
