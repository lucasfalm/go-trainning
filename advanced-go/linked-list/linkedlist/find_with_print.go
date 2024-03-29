package linkedlist

import "fmt"

// NOTE: search among the nodes and print result
func (l *LinkedList) FindWithPrint(value any) (*Node, bool) {
	result, ok := l.Find(value)

	msg := ""

	if ok {
		msg += fmt.Sprintf("found node: %v", result.Value)

		if result.Head != nil {
			msg += fmt.Sprintf(" | head: %v", result.Head.Value)
		}

		if result.Tail != nil {
			msg += fmt.Sprintf(" | tail: %v", result.Tail.Value)
		}
	}

	if !ok {
		msg += "node has not been found"
	}

	fmt.Println(msg)

	return result, ok
}
