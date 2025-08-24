package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	i := 2
	if i >= len(s) || i < 0 {
		return
	}

	copy(s[i:], s[i+1:])
	s[len(s)-1] = 0
	s = s[:len(s)-1]

	fmt.Println(s)
}
