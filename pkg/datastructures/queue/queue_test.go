package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueueEnqueue(t *testing.T) {
	q := NewQueue()
	err := q.Enqueue(1)
	assert.NoError(t, err)
	err = q.Enqueue(2)
	assert.NoError(t, err)
	assert.Equal(t, 2, q.Size(), "Queue size should be 2")
	q.Print()
}

func TestQueueEnqueueFull(t *testing.T) {
	q := NewQueue()
	err := q.Enqueue(1)
	assert.NoError(t, err)
	err = q.Enqueue(2)
	assert.NoError(t, err)
	err = q.Enqueue(3)
	assert.NoError(t, err)
	err = q.Enqueue(4)
	assert.NoError(t, err)
	err = q.Enqueue(5)
	assert.NoError(t, err)
	err = q.Enqueue(6)
	assert.NoError(t, err)
	err = q.Enqueue(7)
	assert.NoError(t, err)
	err = q.Enqueue(8)
	assert.NoError(t, err)
	err = q.Enqueue(9)
	assert.NoError(t, err)
	err = q.Enqueue(10)
	assert.NoError(t, err)
	assert.Equal(t, 10, q.Size(), "Queue size should be 10")
	q.Print()
}

func TestQueueEnqueueOverflow(t *testing.T) {
	q := NewQueue()
	err := q.Enqueue(1)
	assert.NoError(t, err)
	err = q.Enqueue(2)
	assert.NoError(t, err)
	err = q.Enqueue(3)
	assert.NoError(t, err)
	err = q.Enqueue(4)
	assert.NoError(t, err)
	err = q.Enqueue(5)
	assert.NoError(t, err)
	err = q.Enqueue(6)
	assert.NoError(t, err)
	err = q.Enqueue(7)
	assert.NoError(t, err)
	err = q.Enqueue(8)
	assert.NoError(t, err)
	err = q.Enqueue(9)
	assert.NoError(t, err)
	err = q.Enqueue(10)
	assert.NoError(t, err)
	err = q.Enqueue(11)
	assert.Error(t, err, "Queue overflow!")
	q.Print()
}

func TestQueueRoundRobin(t *testing.T) {
	q := NewQueue()
	err := q.Enqueue(1)
	assert.NoError(t, err)
	err = q.Enqueue(2)
	assert.NoError(t, err)
	err = q.Enqueue(3)
	assert.NoError(t, err)
	err = q.Enqueue(4)
	assert.NoError(t, err)
	err = q.Enqueue(5)
	assert.NoError(t, err)
	err = q.Enqueue(6)
	assert.NoError(t, err)
	err = q.Enqueue(7)
	assert.NoError(t, err)
	err = q.Enqueue(8)
	assert.NoError(t, err)
	err = q.Enqueue(9)
	assert.NoError(t, err)
	err = q.Enqueue(10)
	assert.NoError(t, err)
	_, err = q.Dequeue()
	assert.NoError(t, err)
	err = q.Enqueue(11)
	assert.NoError(t, err)
	assert.Equal(t, 10, q.Size(), "Queue size should be 10")
	q.Print()
}

func TestQueueDequeue(t *testing.T) {
	q := NewQueue()
	err := q.Enqueue(1)
	assert.NoError(t, err)
	err = q.Enqueue(2)
	assert.NoError(t, err)
	err = q.Enqueue(3)
	assert.NoError(t, err)

	element, err := q.Dequeue()
	assert.NoError(t, err, "Dequeue should not return an error")
	assert.Equal(t, 1, element, "Dequeue should return 1")
	q.Print()
}

func TestQueuePeekEmpty(t *testing.T) {
	q := NewQueue()
	_, err := q.Peek()
	assert.Error(t, err, "Peek should return an error")
}

func TestQueuePeek(t *testing.T) {
	q := NewQueue()
	err := q.Enqueue(1)
	assert.NoError(t, err)
	err = q.Enqueue(2)
	assert.NoError(t, err)

	element, err := q.Peek()
	assert.NoError(t, err, "Peek should not return an error")
	assert.Equal(t, 1, element, "Peek should return 1")
	q.Print()
}

func TestQueueIsEmptyWithNoEmpty(t *testing.T) {
	q := NewQueue()
	err := q.Enqueue(1)
	assert.NoError(t, err)
	assert.False(t, q.IsEmpty(), "Queue should not be empty")
}

func TestQueueIsEmpty(t *testing.T) {
	q := NewQueue()
	assert.True(t, q.IsEmpty(), "Queue should be empty")
}
