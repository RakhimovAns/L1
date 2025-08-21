package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите время для таймаута: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	n, err := strconv.Atoi(input)
	if err != nil || n <= 0 {
		fmt.Println("Некорректное число")
		return
	}

	ctx, stop := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer stop()

	channel := make(chan int, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Reader done")
				return
			case val := <-channel:
				fmt.Println("value", val)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		value := 1
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Writer done")
				return
			case <-ticker.C:
				channel <- value
				value++
			}
		}
	}()

	wg.Wait()
	fmt.Println("Main done")
}
