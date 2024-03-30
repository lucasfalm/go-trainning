package doubly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	t.Run("a string object", func(t *testing.T) {
		dll.Push("string")

		assert.Equal(t, dll.Any(), true)
	})

	t.Run("a nil object", func(t *testing.T) {
		dll.Push(nil)

		assert.Equal(t, dll.Any(), true)
	})
}
