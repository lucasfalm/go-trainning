package doubly

import (
	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist/node"
)

// NOTE: search among the nodes
func (l *DoublyLinkedList) Find(value any) (node.NodeInterface, bool) {
	if l.Any() {
		for _, node := range l.nodes {
			if node.Value() == value {
				return node, true
			}
		}
	}

	return nil, false
}
