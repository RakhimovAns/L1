package main

import (
	"fmt"
)

func setBit(x int64, k uint, value int) int64 {
	if value == 1 {
		return x | (1 << k)
	}

	return x &^ (1 << k)
}

func main() {
	var num int64 = 5

	num = setBit(num, 0, 0) //---> 4
	fmt.Println(num)

	//num = setBit(num, 2, 1)
	//fmt.Println(num)
}
