package linkedlist

// NOTE: insert the node at the position
func (l *LinkedList) Insert(value any, position int) *Node {
	newNode := &Node{Value: value}

	if l.Any() {
		if l.Count >= position-1 {
			oldNode := l.Nodes[position-1]

			newNode.Head = oldNode.Head
			newNode.Tail = oldNode
			oldNode.Head = newNode

			oldNodes := l.Nodes[position-1:]

			l.Nodes = append(l.Nodes[:position], oldNodes...)

			l.Nodes[position-1] = newNode

		} else {
			l.Nodes = append(l.Nodes, newNode)
		}
	} else {
		l.Nodes = append(l.Nodes, newNode)
	}

	l.Count++

	return newNode
}
