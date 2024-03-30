package doubly

import (
	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist/node"
)

// NOTE: shows the last node
func (l *DoublyLinkedList) Last() node.NodeInterface {
	if l.Any() {
		return l.nodes[len(l.nodes)-1]
	}

	return nil
}
