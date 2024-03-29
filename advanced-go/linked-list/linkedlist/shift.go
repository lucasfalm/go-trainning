package linkedlist

// NOTE: remove the first node
func (l *LinkedList) Shift() *Node {
	if l.Any() {
		firstNode := l.First()

		if firstNode.Tail != nil {
			firstNode.Tail.Head = nil
			firstNode.Tail = nil
		}

		l.count--

		l.nodes = l.nodes[1:]

		return firstNode
	}

	return nil
}
