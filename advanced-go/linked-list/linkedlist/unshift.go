package linkedlist

// NOTE: insert as the first node
func (l *LinkedList) Unshift(value any) NodeInterface {
	newNode := NewNode()

	newNode.SetValue(value)

	if l.Any() {
		previousFirst := l.First()

		previousFirst.SetHead(newNode)
		newNode.SetTail(previousFirst)
	}

	l.nodes = append([]NodeInterface{newNode}, l.nodes...)

	l.count++

	return newNode
}
