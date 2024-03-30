package doubly

import (
	"testing"

	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist/node"
	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	t.Run("a string object", func(t *testing.T) {
		assert.Equal(t, dll.Any(), false)

		assert.Implements(t, (*node.NodeInterface)(nil), dll.Push("string"))

		assert.Equal(t, dll.Any(), true)
	})

	t.Run("a nil object", func(t *testing.T) {
		assert.Equal(t, dll.Any(), false)

		assert.Implements(t, (*node.NodeInterface)(nil), dll.Push(nil))

		assert.Equal(t, dll.Any(), true)
	})
}
