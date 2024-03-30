package main

import (
	"fmt"

	structuresFactory "github.com/lucasfalm/go-training/advanced-go/structures/factory"
)

var (
	ll = structuresFactory.DoublyLinkedList()
)

func main() {
	// examplesOne()

	examplesTwo()
}

func examplesOne() {
	fmt.Println("pushing 32")
	ll.Push(32)
	ll.Print()

	fmt.Println("pushing 10")
	ll.Push(10)
	ll.Print()

	fmt.Println("pushing 5")
	ll.Push(5)
	ll.Print()

	fmt.Println("popping")
	ll.Pop()
	ll.Print()

	fmt.Println("popping")
	ll.Pop()
	ll.Print()

	fmt.Println("popping")
	ll.Pop()
	ll.Print()

	fmt.Println("populating linked list...")
	ll.Push(1)
	ll.Push(290)
	ll.Push(948)
	ll.Push(1920)
	ll.Push(100000)
	ll.Push(2)
	ll.Print()

	fmt.Println("unshifting 88...")
	ll.Unshift(88)
	ll.Print()

	fmt.Println("searching value 1920...")
	ll.FindWithPrint(1920)

	fmt.Println("searching value 5...")
	ll.FindWithPrint(5)

	fmt.Println("shifting...")
	ll.Shift()
	ll.Print()

	fmt.Println("shifting...")
	ll.Shift()
	ll.Print()

	fmt.Println("shifting...")
	ll.Shift()
	ll.Print()

	fmt.Println("inserting at 2 value 19...")
	ll.Insert(19, 2)
	ll.Print()

	fmt.Println("inserting at 1 value 33...")
	ll.Insert(33, 1)
	ll.Print()

	fmt.Println("inserting at 99 value 999...")
	ll.Insert(999, 99)
	ll.Print()
}

func examplesTwo() {
	fmt.Println("populating linked list...")
	ll.Push(1)
	ll.Push(2)
	ll.Push(3)
	ll.Push(4)
	ll.Push(5)
	ll.Push(6)
	ll.Print()

	n, ok := ll.Find(1)
	if ok {
		n.Print()
	}
}
