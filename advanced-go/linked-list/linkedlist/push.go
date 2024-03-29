package linkedlist

// NOTE: insert as the last node
func (l *LinkedList) Push(value any) NodeInterface {
	newNode := NewNode()

	newNode.SetValue(value)

	if l.Any() {
		newNode.SetHead(l.Last())

		l.Last().SetTail(newNode)
	}

	l.nodes = append(l.nodes, newNode)

	l.count++

	return newNode
}
