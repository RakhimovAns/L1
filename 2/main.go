// L1.2
package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}

	arr := [5]int{2, 4, 6, 8, 10}

	wg.Add(len(arr))

	for i := 0; i < len(arr); i++ {
		go func(i int) {
			defer wg.Done()

			mx.Lock()
			arr[i] = arr[i] * arr[i]
			mx.Unlock()

		}(i)
	}

	wg.Wait()

	fmt.Println(arr)
}
