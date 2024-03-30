package main

import (
	"fmt"

	"github.com/lucasfalm/go-training/advanced-go/concurrency/synchronize"
)

func main() {
	fmt.Println("non-sync:")
	synchronize.NonSynchronized(1000)

	fmt.Println("---- x ----- x -----")

	fmt.Println("using mutex:")
	synchronize.UsingMutex(1000)

	fmt.Println("---- x ----- x -----")

	fmt.Println("using channels:")
	synchronize.UsingChannels(1000)
}
