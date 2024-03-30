package doubly

import (
	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist/node"
)

// NOTE: insert as the first node
func (l *DoublyLinkedList) Unshift(value any) node.NodeInterface {
	newNode := node.NewNode()

	newNode.SetValue(value)

	if l.Any() {
		previousFirst := l.First()

		previousFirst.SetHead(newNode)
		newNode.SetTail(previousFirst)
	}

	l.nodes = append([]node.NodeInterface{newNode}, l.nodes...)

	l.count++

	return newNode
}
