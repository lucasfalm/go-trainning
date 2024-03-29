package doubly

import "github.com/lucasfalm/go-training/advanced-go/linked-list/linkedlist"

// NOTE: insert as the last node
func (l *DoublyLinkedList) Push(value any) linkedlist.NodeInterface {
	newNode := linkedlist.NewNode()

	newNode.SetValue(value)

	if l.Any() {
		newNode.SetHead(l.Last())

		l.Last().SetTail(newNode)
	}

	l.nodes = append(l.nodes, newNode)

	l.count++

	return newNode
}
