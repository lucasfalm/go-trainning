package synchronize

import (
	"fmt"
	"sync"
)

func UsingMutex(countTo int) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	counter := 0

	for i := 1; i <= countTo; i++ {
		wg.Add(1)

		go addMutex(&counter, &wg, &mu)
	}

	wg.Wait()

	fmt.Printf("total count: %v\n", counter)
}

func addMutex(counter *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer mu.Unlock()
	defer wg.Done()

	mu.Lock()
	*counter++

	// fmt.Printf("counter %v\n", *counter)
}
