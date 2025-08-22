package main

import (
	"fmt"
	"time"
)

// остановка через канал уведов
func main() {
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Горутина завершена (канал).")
				return
			default:
				fmt.Println("Работаю...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	quit <- struct{}{}
	time.Sleep(1 * time.Second)
}
