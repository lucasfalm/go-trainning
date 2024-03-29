package linkedlist

// NOTE: insert as the last node
func (l *LinkedList) Push(value any) *Node {
	newNode := Node{Value: value}

	if l.Any() {
		newNode.Head = l.Last()

		l.Last().Tail = &newNode
	}

	l.nodes = append(l.nodes, &newNode)

	l.count++

	return &newNode
}
