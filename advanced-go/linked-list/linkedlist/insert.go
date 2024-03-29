package linkedlist

// NOTE: insert the node at the position
func (l *LinkedList) Insert(value any, position int) NodeInterface {
	newNode := NewNode()

	newNode.SetValue(value)

	if l.Any() {
		if l.count >= position-1 {
			oldNode := l.nodes[position-1]

			newNode.SetHead(oldNode.Head())
			newNode.SetTail(oldNode)
			oldNode.SetHead(newNode)

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
