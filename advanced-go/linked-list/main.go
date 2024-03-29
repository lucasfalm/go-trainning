package main

import (
	"fmt"

	"github.com/lucasfalm/go-training/advanced-go/linked-list/linkedlist/doubly"
)

var (
	linkedList = doubly.NewLinkedList()
)

func main() {
	// examplesOne()

	examplesTwo()
}

func examplesOne() {
	fmt.Println("pushing 32")
	linkedList.Push(32)
	linkedList.Print()

	fmt.Println("pushing 10")
	linkedList.Push(10)
	linkedList.Print()

	fmt.Println("pushing 5")
	linkedList.Push(5)
	linkedList.Print()

	fmt.Println("popping")
	linkedList.Pop()
	linkedList.Print()

	fmt.Println("popping")
	linkedList.Pop()
	linkedList.Print()

	fmt.Println("popping")
	linkedList.Pop()
	linkedList.Print()

	fmt.Println("populating linked list...")
	linkedList.Push(1)
	linkedList.Push(290)
	linkedList.Push(948)
	linkedList.Push(1920)
	linkedList.Push(100000)
	linkedList.Push(2)
	linkedList.Print()

	fmt.Println("unshifting 88...")
	linkedList.Unshift(88)
	linkedList.Print()

	fmt.Println("searching value 1920...")
	linkedList.FindWithPrint(1920)

	fmt.Println("searching value 5...")
	linkedList.FindWithPrint(5)

	fmt.Println("shifting...")
	linkedList.Shift()
	linkedList.Print()

	fmt.Println("shifting...")
	linkedList.Shift()
	linkedList.Print()

	fmt.Println("shifting...")
	linkedList.Shift()
	linkedList.Print()

	fmt.Println("inserting at 2 value 19...")
	linkedList.Insert(19, 2)
	linkedList.Print()

	fmt.Println("inserting at 1 value 33...")
	linkedList.Insert(33, 1)
	linkedList.Print()

	fmt.Println("inserting at 99 value 999...")
	linkedList.Insert(999, 99)
	linkedList.Print()
}

func examplesTwo() {
	fmt.Println("populating linked list...")
	linkedList.Push(1)
	linkedList.Push(2)
	linkedList.Push(3)
	linkedList.Push(4)
	linkedList.Push(5)
	linkedList.Push(6)
	linkedList.Print()

	n, ok := linkedList.Find(1)
	if ok {
		n.Print()
	}
}
