package linkedlist

// NOTE: shows the last node
func (l *LinkedList) Last() NodeInterface {
	if l.Any() {
		return l.nodes[len(l.nodes)-1]
	}

	return nil
}
