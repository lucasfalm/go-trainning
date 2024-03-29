package linkedlist

// NOTE: remove the last node
func (l *LinkedList) Pop() NodeInterface {
	if l.Any() {
		poppedNode := l.nodes[l.count-1]

		if poppedNode.Head() != nil {
			poppedNode.Head().SetTail(nil)
			poppedNode.SetHead(nil)
		}

		l.count--

		l.nodes = l.nodes[:l.count]

		return poppedNode
	}

	return nil
}
