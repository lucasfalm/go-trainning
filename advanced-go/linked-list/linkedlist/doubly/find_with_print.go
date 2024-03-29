package doubly

import (
	"fmt"

	"github.com/lucasfalm/go-training/advanced-go/linked-list/linkedlist"
)

// NOTE: search among the nodes and print result
func (l *DoublyLinkedList) FindWithPrint(value any) (linkedlist.NodeInterface, bool) {
	result, ok := l.Find(value)

	msg := ""

	if ok {
		msg += fmt.Sprintf("found node: %v", result.Value())

		if result.Head() != nil {
			msg += fmt.Sprintf(" | head: %v", result.Head().Value())
		}

		if result.Tail() != nil {
			msg += fmt.Sprintf(" | tail: %v", result.Tail().Value())
		}
	}

	if !ok {
		msg += "node has not been found"
	}

	fmt.Println(msg)

	return result, ok
}
