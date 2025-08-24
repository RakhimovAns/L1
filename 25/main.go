package main

import (
	"fmt"
	"time"
)

func mySleep(duration time.Duration) {
	//реализация фором
	//start := time.Now()
	//for time.Since(start) < duration {
	//}

	//через канал
	done := make(chan struct{})
	time.AfterFunc(duration, func() {
		close(done)
	})
	<-done
}

func main() {
	fmt.Println("Start")
	mySleep(2 * time.Second)
	fmt.Println("End")
}
