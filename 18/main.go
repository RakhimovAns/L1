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

type Counter struct {
	mx    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mx.Lock()
	c.value++
	c.mx.Unlock()
}

func (c *Counter) Value() int {
	c.mx.Lock()
	defer c.mx.Unlock()
	return c.value
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	wg := sync.WaitGroup{}
	counter := &Counter{}

	worker := func() {
		defer wg.Done()
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				counter.Inc()
			}
		}
	}

	wg.Add(2)
	go worker()
	go worker()

	wg.Wait()

	fmt.Println(counter.Value())
}
