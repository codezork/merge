package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Len(t *testing.T) {
	t.Run("stack len", func(t *testing.T) {
		stack := new(Stack)
		stack.Push("1")
		stack.Push("2")
		stack.Push("3")
		assert.Equal(t, 3, stack.Len(), "stack size invalid")
	})
}

func TestStack_Pop(t *testing.T) {
	t.Run("stack pop", func(t *testing.T) {
		stack := new(Stack)
		stack.Push("1")
		stack.Push("2")
		stack.Push("3")

		assert.Equal(t, "3", stack.Pop(), "stack pop failed")
	})
}

func TestStack_Push(t *testing.T) {
	t.Run("stack pop", func(t *testing.T) {
		stack := new(Stack)
		stack.Push("1")
		stack.Push("2")
		stack.Push("3")

		assert.Equal(t, "3", stack.Top(), "stack push failed")
	})
}

func TestStack_Top(t *testing.T) {
	t.Run("stack pop", func(t *testing.T) {
		stack := new(Stack)
		stack.Push("1")
		stack.Push("2")
		stack.Push("3")
		stack.Pop()

		assert.Equal(t, "2", stack.Top(), "stack top invalid")
	})
}