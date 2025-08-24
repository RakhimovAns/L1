package main

import (
	"fmt"
	"strings"
)

func uniqueChars(s string) bool {
	s = strings.ToLower(s)
	seen := make(map[rune]bool)

	for _, c := range s {
		if seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}

func main() {
	fmt.Println(uniqueChars("abcd"))
	fmt.Println(uniqueChars("abCdefAaf"))
	fmt.Println(uniqueChars("aabcd"))
	fmt.Println(uniqueChars("Привет"))
}
