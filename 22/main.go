package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1_500_000)
	b := big.NewInt(2_000_000)

	sum := new(big.Int).Add(a, b)
	fmt.Println("sum:", sum)

	diff := new(big.Int).Sub(a, b)
	fmt.Println("diff:", diff)

	prod := new(big.Int).Mul(a, b)
	fmt.Println("prod:", prod)

	if b.Cmp(big.NewInt(0)) != 0 {
		quot := new(big.Int).Div(a, b)
		fmt.Println("quot:", quot)
	} else {
		fmt.Println("error: division by zero")
	}
}
