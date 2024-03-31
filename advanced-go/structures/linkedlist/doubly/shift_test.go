package doubly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShift(t *testing.T) {
	t.Run("when there is no node", func(t *testing.T) {
		setup()

		result := dll.Shift()

		assert.Nil(t, result)
	})

	t.Run("when there is one node", func(t *testing.T) {
		setup()

		n := dll.Push(1)

		assert.True(t, dll.Any())

		result := dll.Shift()

		assert.False(t, dll.Any())

		assert.Equal(t, n, result)
	})

	t.Run("when there are multiple nodes", func(t *testing.T) {
		setup()

		n1 := dll.Push(1)
		n2 := dll.Push(2)

		result := dll.Shift()

		assert.Equal(t, n1, result)

		assert.Nil(t, n1.Head())

		assert.Nil(t, n1.Tail())

		assert.Nil(t, n2.Head())

		assert.Nil(t, n2.Tail())
	})
}
