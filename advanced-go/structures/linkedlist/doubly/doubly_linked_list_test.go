package doubly

import (
	"os"
	"testing"

	"github.com/lucasfalm/go-training/advanced-go/structures/linkedlist"
	"github.com/stretchr/testify/assert"
)

var (
	dll linkedlist.LinkedListInterface
)

func setup() {
	dll = &DoublyLinkedList{}
}

func teardown() {}

func withSetup(t *testing.T, testFunc func(t *testing.T)) {
	setup()

	defer teardown()

	testFunc(t)
}

func TestMain(m *testing.M) {
	// Call setup function before running tests
	setup()

	// Run tests
	exitCode := m.Run()

	// Perform teardown actions here if needed
	teardown()

	// Exit with the same exit code as the tests
	// This ensures proper exit code propagation
	os.Exit(exitCode)
}

func TestNewLinkedList(t *testing.T) {
	result := NewLinkedList()

	assert.Implements(t, (*linkedlist.LinkedListInterface)(nil), result)

	assert.NotNil(t, result)
}
