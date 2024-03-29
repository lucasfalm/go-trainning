package linkedlist

// NOTE: insert as the first node
func (l *LinkedList) Unshift(value any) *Node {
	newNode := &Node{Value: value}

	if l.Any() {
		previousFirst := l.First()

		previousFirst.Head = newNode
		newNode.Tail = previousFirst
	}

	l.nodes = append([]*Node{newNode}, l.nodes...)

	l.count++

	return newNode
}
