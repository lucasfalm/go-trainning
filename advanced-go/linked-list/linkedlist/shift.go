package linkedlist

// NOTE: remove the first node
func (l *LinkedList) Shift() NodeInterface {
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
