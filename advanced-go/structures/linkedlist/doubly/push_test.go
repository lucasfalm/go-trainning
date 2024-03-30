package doubly

import (
	"testing"

	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist/node"
	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	t.Run("a string object", func(t *testing.T) {
		setup()

		assert.False(t, dll.Any())

		result := dll.Push("string")

		assert.Implements(t, (*node.NodeInterface)(nil), result)

		assert.True(t, dll.Any())

		assert.Nil(t, result.Head())

		assert.Nil(t, result.Tail())

		assert.Equal(t, "string", result.Value())
	})

	t.Run("a nil object", func(t *testing.T) {
		withSetup(t, func(t *testing.T) {
			assert.False(t, dll.Any())

			result := dll.Push(nil)

			assert.Implements(t, (*node.NodeInterface)(nil), result)

			assert.True(t, dll.Any())

			assert.Nil(t, result.Head())

			assert.Nil(t, result.Tail())

			assert.Nil(t, result.Value())
		})
	})

	t.Run("multiple pushes", func(t *testing.T) {
		setup()

		assert.Equal(t, false, dll.Any())

		// NOTE: first push
		resultOne := dll.Push(1)

		assert.Implements(t, (*node.NodeInterface)(nil), resultOne)

		assert.True(t, dll.Any())

		assert.Nil(t, resultOne.Head())

		assert.Nil(t, resultOne.Tail())

		assert.Equal(t, 1, resultOne.Value())

		// NOTE: second push
		resultTwo := dll.Push(2)

		assert.Implements(t, (*node.NodeInterface)(nil), resultTwo)

		assert.Equal(t, resultOne, resultTwo.Head())

		assert.Equal(t, resultTwo, resultOne.Tail())

		assert.Nil(t, resultTwo.Tail())

		assert.Equal(t, 2, resultTwo.Value())

		// NOTE: third push
		resultThree := dll.Push(3)

		assert.Implements(t, (*node.NodeInterface)(nil), resultThree)

		assert.Equal(t, resultTwo, resultThree.Head())

		assert.Nil(t, resultThree.Tail())

		assert.Equal(t, resultThree, resultTwo.Tail())

		assert.Equal(t, 3, resultThree.Value())
	})
}
