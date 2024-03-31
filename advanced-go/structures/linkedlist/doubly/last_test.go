package doubly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLast(t *testing.T) {
	t.Run("when there is no node", func(t *testing.T) {
		setup()

		result := dll.Last()

		assert.Nil(t, result)
	})

	t.Run("when there is at least one node", func(t *testing.T) {
		setup()

		dll.Push("hello")

		result := dll.Last()

		assert.NotNil(t, result)

		assert.Equal(t, "hello", result.Value())
	})
}
