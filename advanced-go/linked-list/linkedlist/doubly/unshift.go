package doubly

import "github.com/lucasfalm/go-training/advanced-go/linked-list/linkedlist"

// NOTE: insert as the first node
func (l *DoublyLinkedList) Unshift(value any) linkedlist.NodeInterface {
	newNode := linkedlist.NewNode()

	newNode.SetValue(value)

	if l.Any() {
		previousFirst := l.First()

		previousFirst.SetHead(newNode)
		newNode.SetTail(previousFirst)
	}

	l.nodes = append([]linkedlist.NodeInterface{newNode}, l.nodes...)

	l.count++

	return newNode
}
