package doubly

import (
	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist/node"
)

// NOTE: insert the node at the position
func (l *DoublyLinkedList) Insert(value any, position int) node.NodeInterface {
	newNode := node.NewNode()

	newNode.SetValue(value)

	if l.hasPosition(position) {
		oldNode := l.nodes[position-1]

		newNode.SetHead(oldNode.Head())
		newNode.SetTail(oldNode)
		oldNode.SetHead(newNode)

		oldNodes := l.nodes[position-1:]

		l.nodes = append(l.nodes[:position], oldNodes...)

		l.nodes[position-1] = newNode
	} else {
		l.nodes = append(l.nodes, newNode)
	}

	l.count++

	return newNode
}

func (l *DoublyLinkedList) hasPosition(position int) bool {
	return l.count >= position-1
}
