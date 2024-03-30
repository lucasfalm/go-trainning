package doubly

import (
	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist/node"
)

// NOTE: remove the first node
func (l *DoublyLinkedList) Shift() node.NodeInterface {
	if l.Any() {
		firstNode := l.First()

		if firstNode.Tail() != nil {
			firstNode.Tail().SetHead(nil)
			firstNode.SetTail(nil)
		}

		l.count--

		l.nodes = l.nodes[1:]

		return firstNode
	}

	return nil
}
