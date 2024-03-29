package doubly

import "github.com/lucasfalm/go-training/advanced-go/linked-list/linkedlist"

// NOTE: shows the last node
func (l *DoublyLinkedList) Last() linkedlist.NodeInterface {
	if l.Any() {
		return l.nodes[len(l.nodes)-1]
	}

	return nil
}
