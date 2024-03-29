package doubly

import "github.com/lucasfalm/go-training/advanced-go/linked-list/linkedlist"

// NOTE: doubly linked list
type DoublyLinkedList struct {
	nodes []linkedlist.NodeInterface
	count int
}

// NOTE: creates a double linked list and returns the interface/contract
func NewLinkedList() linkedlist.LinkedListInterface {
	return &DoublyLinkedList{}
}
