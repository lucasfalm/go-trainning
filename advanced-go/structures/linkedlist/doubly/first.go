package doubly

import (
	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist/node"
)

// NOTE: shows the first node
func (l *DoublyLinkedList) First() node.NodeInterface {
	if l.Any() {
		return l.nodes[0]
	}

	return nil
}
