package doubly

import "github.com/lucasfalm/go-training/advanced-go/linked-list/linkedlist"

// NOTE: remove the first node
func (l *DoublyLinkedList) Shift() linkedlist.NodeInterface {
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
