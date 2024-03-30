package doubly

import (
	"testing"

	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist/node"
	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	t.Run("when there is no node", func(t *testing.T) {
		setup()

		result := dll.First()

		assert.Equal(t, nil, result)
	})

	t.Run("when there are nodes", func(t *testing.T) {
		setup()

		dll.Push(nil)

		result := dll.First()

		assert.Implements(t, (*node.NodeInterface)(nil), result)
	})
}
