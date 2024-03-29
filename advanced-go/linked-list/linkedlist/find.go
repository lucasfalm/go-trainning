package linkedlist

// NOTE: search among the nodes
func (l *LinkedList) Find(value any) (*Node, bool) {
	if l.Any() {
		for _, node := range l.nodes {
			if node.Value == value {
				return node, true
			}
		}
	}

	return nil, false
}
