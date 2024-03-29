package linkedlist

// NOTE: insert the node at the position
func (l *LinkedList) Insert(value any, position int) *Node {
	newNode := &Node{Value: value}

	if l.Any() {
		if l.count >= position-1 {
			oldNode := l.nodes[position-1]

			newNode.Head = oldNode.Head
			newNode.Tail = oldNode
			oldNode.Head = newNode

			oldNodes := l.nodes[position-1:]

			l.nodes = append(l.nodes[:position], oldNodes...)

			l.nodes[position-1] = newNode

		} else {
			l.nodes = append(l.nodes, newNode)
		}
	} else {
		l.nodes = append(l.nodes, newNode)
	}

	l.count++

	return newNode
}
