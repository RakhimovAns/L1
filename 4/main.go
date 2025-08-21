package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Break the loop")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Hello in a loop")
			}
		}
	}()

	wg.Wait()
	fmt.Println("Main done")
}

//Обоснование данного метода
//в программе использую signal.NotifyContext, который создаёт контекст, автоматически отменяемый при получении сигналов прерывания (SIGINT, SIGTERM), что позволяет не блокировать главную горутину.
//Внутри горутины <-ctx.Done обработка завершения
//Испольую waitGroup, чтоб ждать завершение всех горутин
