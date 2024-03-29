package doubly

import "github.com/lucasfalm/go-training/advanced-go/linked-list/linkedlist"

// NOTE: search among the nodes
func (l *DoublyLinkedList) Find(value any) (linkedlist.NodeInterface, bool) {
	if l.Any() {
		for _, node := range l.nodes {
			if node.Value() == value {
				return node, true
			}
		}
	}

	return nil, false
}
