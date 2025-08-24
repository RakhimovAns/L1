package main

import (
	"fmt"
)

func main() {
	s := []byte("dep and dodep")
	n := len(s)

	reverse := func(b []byte, i, j int) {
		for i < j {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}
	}

	reverse(s, 0, n-1)

	start := 0
	for start < n {
		end := start
		for end < n && s[end] != ' ' {
			end++
		}
		reverse(s, start, end-1)
		start = end + 1
	}

	fmt.Println(string(s))
}
