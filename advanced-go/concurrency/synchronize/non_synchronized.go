package synchronize

import (
	"fmt"
	"sync"
)

func NonSynchronized(countTo int) {
	wg := &sync.WaitGroup{}

	counter := 0

	for i := 1; i <= countTo; i++ {
		wg.Add(1)

		go addNonSynchronized(&counter, wg)
	}

	fmt.Printf("total counter: %v\n", counter)
}

func addNonSynchronized(counter *int, wg *sync.WaitGroup) {
	defer wg.Done()

	*counter++

	// fmt.Printf("counter %v\n", *counter)
}
