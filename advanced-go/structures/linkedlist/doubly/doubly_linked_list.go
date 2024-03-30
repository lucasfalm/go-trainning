package doubly

import (
	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist"
	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist/node"
)

// NOTE: doubly linked list
type DoublyLinkedList struct {
	nodes []node.NodeInterface
	count int
}

// NOTE: creates a double linked list and returns the interface/contract
func NewLinkedList() linkedlist.LinkedListInterface {
	return &DoublyLinkedList{}
}
