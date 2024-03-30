package doubly

import (
	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist/node"
)

// NOTE: remove the last node
func (l *DoublyLinkedList) Pop() node.NodeInterface {
	if l.Any() {
		poppedNode := l.nodes[l.count-1]

		if poppedNode.Head() != nil {
			poppedNode.Head().SetTail(nil)
			poppedNode.SetHead(nil)
		}

		l.count--

		l.nodes = l.nodes[:l.count]

		return poppedNode
	}

	return nil
}
