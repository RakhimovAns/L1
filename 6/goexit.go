package main

import (
	"fmt"
	"runtime"
	"time"
)

// runtime.Goexit()
func main() {
	go func() {
		defer fmt.Println("defer сработал перед выходом.")
		fmt.Println("Работаю...")
		time.Sleep(1 * time.Second)
		runtime.Goexit()
		fmt.Println("Эта строка не выполнится.")
	}()

	time.Sleep(2 * time.Second)
}
