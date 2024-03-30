package doubly

import (
	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist/node"
)

// NOTE: insert the node at the position
func (l *DoublyLinkedList) Insert(value any, position int) node.NodeInterface {
	newNode := node.NewNode()

	newNode.SetValue(value)

	if l.hasPosition(position) {
		if position > 0 {
			position = position - 1
		}

		oldNode := l.nodes[position]

		newNode.SetHead(oldNode.Head())
		newNode.SetTail(oldNode)
		oldNode.SetHead(newNode)

		oldNodes := l.nodes[position:]

		l.nodes = append(l.nodes[:position], oldNodes...)

		l.nodes[position] = newNode
	} else {
		lastNode := l.Last()

		if lastNode != nil {
			lastNode.SetTail(newNode)
			newNode.SetHead(lastNode)
		}

		l.nodes = append(l.nodes, newNode)
	}

	l.count++

	return newNode
}

func (l *DoublyLinkedList) hasPosition(position int) bool {
	return l.count >= position
}
