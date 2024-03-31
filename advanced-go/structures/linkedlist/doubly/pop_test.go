package doubly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPop(t *testing.T) {
	t.Run("when there is no node", func(t *testing.T) {
		setup()

		result := dll.Pop()

		assert.Nil(t, result)
	})

	t.Run("when there is one node", func(t *testing.T) {
		setup()

		n := dll.Push(1)

		assert.True(t, dll.Any())

		result := dll.Pop()

		assert.False(t, dll.Any())

		assert.NotNil(t, result)

		assert.Equal(t, n, result)

		assert.Equal(t, 1, result.Value())

		assert.Nil(t, result.Head())

		assert.Nil(t, result.Tail())
	})

	t.Run("when there are two nodes", func(t *testing.T) {
		setup()

		n1 := dll.Push(1)
		n2 := dll.Push(2)

		assert.Equal(t, n1.Tail(), n2)
		assert.Equal(t, n2.Head(), n1)

		assert.True(t, dll.Any())

		result := dll.Pop()

		assert.True(t, dll.Any())

		assert.NotNil(t, result)

		assert.Equal(t, n2, result)

		assert.Equal(t, 2, result.Value())

		assert.Nil(t, result.Head())

		assert.Nil(t, result.Tail())

		assert.Nil(t, n1.Head())

		assert.Nil(t, n1.Tail())
	})
}
