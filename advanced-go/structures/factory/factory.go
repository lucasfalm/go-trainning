package factory

import (
	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist"
	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist/doubly"
)

func DoublyLinkedList() linkedlist.LinkedListInterface {
	return doubly.NewLinkedList()
}
