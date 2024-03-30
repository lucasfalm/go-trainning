package synchronize

import (
	"fmt"
)

func UsingChannels(countTo int) {
	// NOTE: the buffer tells how many items can be in the channel at the same time
	cBuffer := 1

	ch := make(chan *int, cBuffer)
	defer close(ch)

	counter := 0

	for i := 0; i <= countTo; i++ {
		// NOTE: the channel will block/wait until a new item can be added (max buffer per time)
		// 			 a.k.a: as soon as another goroutine releases/consumes an item from the channel
		ch <- &counter

		go printChannel(ch)
	}

	fmt.Printf("total counter: %v\n", counter)
}

func printChannel(ch chan *int) {
	// NOTE: consuming/releasing an item from the channel
	// 			 allowing to another goroutines add a new item to it
	counter := <-ch

	// NOTE: as this channel has a buffer of one, only one goroutines per time will
	// 			 acquire the channel and change the counter
	*counter++

	// fmt.Printf("counter %v\n", *counter)
}
