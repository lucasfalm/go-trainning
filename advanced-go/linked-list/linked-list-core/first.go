package linked_list_core

// NOTE: shows the first node
func (l *LinkedList) First() *Node {
	if l.Any() {
		return l.Nodes[0]
	}

	return nil
}
