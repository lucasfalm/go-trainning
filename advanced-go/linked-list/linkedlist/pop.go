package linkedlist

// NOTE: remove the last node
func (l *LinkedList) Pop() *Node {
	if l.Any() {
		poppedNode := l.nodes[l.count-1]

		if poppedNode.Head != nil {
			poppedNode.Head.Tail = nil
			poppedNode.Head = nil
		}

		l.count--

		l.nodes = l.nodes[:l.count]

		return poppedNode
	}

	return nil
}
