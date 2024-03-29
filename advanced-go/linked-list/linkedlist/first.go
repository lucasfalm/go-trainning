package linkedlist

// NOTE: shows the first node
func (l *LinkedList) First() *Node {
	if l.Any() {
		return l.nodes[0]
	}

	return nil
}
