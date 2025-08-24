package main

import "fmt"

func reverse(s string) string {
	temp := []rune(s)

	for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
		temp[i], temp[j] = temp[j], temp[i]

	}

	return string(temp)
}

func main() {
	str := "abcd"
	fmt.Println(reverse(str))
}
