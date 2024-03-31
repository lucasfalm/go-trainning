package doubly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	t.Run("when there is already a node at the position", func(t *testing.T) {
		setup()

		previousNodeAtPosition := dll.Push(2)

		result := dll.Insert(1, 0)

		assert.NotNil(t, result)

		assert.Equal(t, previousNodeAtPosition, result.Tail())

		assert.Equal(t, result, previousNodeAtPosition.Head())

		assert.Equal(t, 1, result.Value())
	})

	t.Run("when there is no node at the position", func(t *testing.T) {
		setup()

		firstNode := dll.Push(1)

		result := dll.Insert(2, 2)

		assert.NotNil(t, result)

		assert.Equal(t, result, firstNode.Tail())

		assert.Equal(t, firstNode, result.Head())

		assert.Equal(t, 2, result.Value())
	})

	t.Run("when there is no node", func(t *testing.T) {
		setup()

		result := dll.Insert(2, 100)

		assert.NotNil(t, result)

		assert.Equal(t, result, dll.First())

		assert.Equal(t, 2, result.Value())
	})
}
