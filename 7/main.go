package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var (
		mx sync.Mutex
		wg sync.WaitGroup
	)

	mp := make(map[string]string)

	numWorkers := 10
	numOps := 100

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOps; j++ {
				key := "key_" + strconv.Itoa(id) + "_" + strconv.Itoa(j)
				val := "value_" + strconv.Itoa(j)

				mx.Lock()
				mp[key] = val
				mx.Unlock()
			}
		}(i)
	}

	wg.Wait()

	fmt.Println("Общее количество записей:", len(mp))
}
