package doubly

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	t.Run("when there is none", func(t *testing.T) {
		assert.Equal(t, ll.Any(), false)
	})

	t.Run("when there is one", func(t *testing.T) {
		ll.Push(1)

		assert.Equal(t, ll.Any(), true)
	})

	t.Run("when there are multiple", func(t *testing.T) {
		for i := 0; i <= 30; i++ {
			ll.Push(i)
		}

		assert.Equal(t, ll.Any(), true)
	})
}
