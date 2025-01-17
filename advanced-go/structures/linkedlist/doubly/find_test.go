package doubly

import (
	"testing"

	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist/node"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	t.Run("when there is no node", func(t *testing.T) {
		setup()

		result, ok := dll.Find(1)

		assert.False(t, ok)
		assert.Nil(t, result)
	})

	t.Run("when there are nodes but none that matches", func(t *testing.T) {
		setup()

		dll.Push(2)

		result, ok := dll.Find(1)

		assert.False(t, ok)
		assert.Nil(t, result)
	})

	t.Run("when there are nodes and a match", func(t *testing.T) {
		setup()

		dll.Push(1)

		result, ok := dll.Find(1)

		assert.True(t, ok)
		assert.Implements(t, (*node.NodeInterface)(nil), result)
	})
}
