package main

import (
	"fmt"
	"time"
)

// Остановка по условию внутри горутины
func main() {
	stop := false

	go func() {
		for {
			if stop {
				fmt.Println("Горутина завершена (условие).")
				return
			}
			fmt.Println("Работаю...")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)
	stop = true
	time.Sleep(1 * time.Second)
}
