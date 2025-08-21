// L1.3
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d got job %d\n", id, job)
	}
}

// запуск через go run main.go
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите количество воркеров: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	n, err := strconv.Atoi(input)
	if err != nil || n <= 0 {
		fmt.Println("Некорректное число")
		return
	}

	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	for i := 1; i <= n; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	go func() {
		counter := 1
		for {
			jobs <- counter
			counter++
			time.Sleep(500 * time.Millisecond)
		}
	}()

	wg.Wait()
}
