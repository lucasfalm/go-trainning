package doubly

import (
	"testing"

	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist/node"
	"github.com/stretchr/testify/assert"
)

// NOTE: replacing the real doubly struct but this mock
type doublyLinkedListMockPush struct {
	DoublyLinkedList
}

func (dll *doublyLinkedListMockPush) Push(value any) node.NodeInterface {
	dll.count++

	return nil
}

func TestAny(t *testing.T) {
	dll = &doublyLinkedListMockPush{}

	t.Run("when there is none", func(t *testing.T) {
		assert.Equal(t, dll.Any(), false)
	})

	t.Run("when there is one", func(t *testing.T) {
		dll.Push(1)

		assert.Equal(t, dll.Any(), true)
	})

	t.Run("when there are multiple", func(t *testing.T) {
		for i := 0; i <= 30; i++ {
			dll.Push(i)
		}

		assert.Equal(t, dll.Any(), true)
	})
}
