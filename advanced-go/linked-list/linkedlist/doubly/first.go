package doubly

import "github.com/lucasfalm/go-training/advanced-go/linked-list/linkedlist"

// NOTE: shows the first node
func (l *DoublyLinkedList) First() linkedlist.NodeInterface {
	if l.Any() {
		return l.nodes[0]
	}

	return nil
}
