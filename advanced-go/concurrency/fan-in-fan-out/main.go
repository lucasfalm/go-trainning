package main

import (
	"fmt"
	"sync"
)

var processTwoPartitionSize int

type Item struct{}

func processOne() int {
	fmt.Println("processing #1")
	fmt.Println("finished processing #1")

	return processTwoPartitionSize
}

func processTwo(qtd int) <-chan *Item {
	fmt.Println("processing #2")

	out := make(chan *Item, qtd)
	defer close(out)

	wg := sync.WaitGroup{}

	for i := 1; i <= qtd; i++ {
		wg.Add(1)

		// NOTE: this will break down the processTwo results into smaller chunks
		// and put each in the out channel (and it will be glued together later)
		go executePartitionProcessTwo(out, &wg)
	}

	wg.Wait()

	fmt.Println("finished processing #2")

	return out
}

func processThree(items []*Item) {
	fmt.Printf("processing #3 with items: %v\n", len(items))
	fmt.Println("finished processing #3")
}

func main() {
	processTwoPartitionSize = 30

	p1r := processOne()

	// NOTE: this is slow, so it has been partitioned
	// now it returns a channel that will receive the results from the
	// several goroutines that worked to make each small parts of it
	in := processTwo(p1r)

	// NOTE: this is glueing the results from the several goroutines that
	// sent a message in the channel
	p2r := processP2rResults(in)

	processThree(p2r)
}

// NOTE: can only insert in the channel (chan<- *Item)
func executePartitionProcessTwo(out chan<- *Item, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("processing #2 - sub process")
	out <- &Item{}
	fmt.Println("finished processing #2 - sub process")
}

// NOTE: can only read from the channel (<- chan *Item)
func processP2rResults(in <-chan *Item) []*Item {
	fmt.Println("started reading sub processes #2 results")

	items := []*Item{}

	for item := range in {
		items = append(items, item)
	}

	fmt.Println("finished reading sub processes #2 results")

	return items
}
