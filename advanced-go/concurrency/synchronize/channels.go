package synchronize

import (
	"fmt"
	"sync"
)

func UsingChannels(countTo int) {
	// NOTE: the buffer tells how many items can be in the channel at the same time
	cBuffer := 1

	ch := make(chan *int, cBuffer)
	defer close(ch)

	wg := sync.WaitGroup{}

	counter := 0

	for i := 1; i <= countTo; i++ {
		// NOTE: the channel will block/wait until a new item can be added (max buffer per time)
		// 			 a.k.a: as soon as another thread releases/consumes an item from the channel
		ch <- &counter

		wg.Add(1)

		go printChannel(&wg, ch)
	}

	wg.Wait()

	fmt.Printf("total counter: %v\n", counter)
}

func printChannel(wg *sync.WaitGroup, ch chan *int) {
	defer wg.Done()

	// NOTE: consuming/releasing an item from the channel
	// 			 allowing to another thread add a new item to it
	counter := <-ch

	// NOTE: as this channel has a buffer of one, only one thread per time will
	// 			 acquire the channel and change the counter
	*counter++

	// fmt.Printf("counter %v\n", *counter)
}
