package main

import (
	"fmt"
	"os"
	"time"
)

// os.Exit()
func main() {
	go func() {
		for {
			fmt.Println("Работаю...")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Завершение программы через os.Exit.")
	os.Exit(0)
}
