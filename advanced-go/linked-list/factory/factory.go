package factory

import (
	"github.com/lucasfalm/go-training/advanced-go/linked-list/linkedlist"
	"github.com/lucasfalm/go-training/advanced-go/linked-list/linkedlist/doubly"
)

func DoublyLinkedList() linkedlist.LinkedListInterface {
	return doubly.NewLinkedList()
}
