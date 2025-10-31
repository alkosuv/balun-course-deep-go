package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func TestCircularQueueWithoutLen(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueueWithoutLen[int](queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, -1, queue.Front())
	assert.Equal(t, -1, queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, queue.getQueue()))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, 1, queue.Front())
	assert.Equal(t, 3, queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int{4, 2, 3}, queue.getQueue()))

	assert.Equal(t, 2, queue.Front())
	assert.Equal(t, 4, queue.Back())

	value, ok := queue.Dequeue()
	assert.Equal(t, 2, value)
	assert.Equal(t, true, ok)

	value, ok = queue.Dequeue()
	assert.Equal(t, 3, value)
	assert.Equal(t, true, ok)

	value, ok = queue.Dequeue()
	assert.Equal(t, 4, value)
	assert.Equal(t, true, ok)

	value, ok = queue.Dequeue()
	assert.Equal(t, -1, value)
	assert.Equal(t, false, ok)

	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
