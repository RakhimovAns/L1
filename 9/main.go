package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}

	first := make(chan int)
	second := make(chan int)

	go func() {
		for _, x := range nums {
			first <- x
		}
		close(first)
	}()

	go func() {
		for x := range first {
			second <- x * 2
		}
		close(second)
	}()

	for res := range second {
		fmt.Println(res)
	}
}
