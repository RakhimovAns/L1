package main

import (
	"fmt"
)

func detectType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Тип: int")
	case string:
		fmt.Println("Тип: string")
	case bool:
		fmt.Println("Тип: bool")
	case chan int, chan string, chan bool, chan interface{}:
		fmt.Println("Тип: chan")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	detectType(42)
	detectType("hello")
	detectType(true)
	detectType(make(chan int))
	detectType(3.14)
}
