package main

type CircularQueueWithoutLen[T int8 | int16 | int32 | int64 | int] struct {
	queue []T
	cap   int
	head  int
	tail  int
}

// NewCircularQueueWithoutLen создает новый кольцевой буфер без длины
func NewCircularQueueWithoutLen[T int8 | int16 | int32 | int64 | int](size int) *CircularQueueWithoutLen[T] {
	return &CircularQueueWithoutLen[T]{
		queue: make([]T, size),
		cap:   size,
		head:  -1,
		tail:  -1,
	}
}

func (q *CircularQueueWithoutLen[T]) Dequeue() (T, bool) {
	return q.Front(), q.Pop()
}

func (q *CircularQueueWithoutLen[T]) Push(value T) bool {
	if q.Full() {
		return false
	}

	if q.head == -1 {
		q.head = 0
	}

	q.tail = (q.tail + 1) % q.cap
	q.queue[q.tail] = value

	return true
}

func (q *CircularQueueWithoutLen[T]) Pop() bool {
	if q.Empty() {
		return false
	}

	if q.head == q.tail {
		q.head, q.tail = -1, -1
	} else {
		q.head = (q.head + 1) % q.cap
	}

	return true
}

func (q *CircularQueueWithoutLen[T]) Front() T {
	if q.Empty() {
		return -1
	}

	return q.queue[q.head]
}

func (q *CircularQueueWithoutLen[T]) Back() T {
	if q.Empty() {
		return -1
	}

	return q.queue[q.tail]
}

func (q *CircularQueueWithoutLen[T]) Full() bool {
	return q.head == (q.tail+1)%q.cap
}

func (q *CircularQueueWithoutLen[T]) Empty() bool {
	return q.head == -1
}

func (q *CircularQueueWithoutLen[T]) getQueue() []T {
	return q.queue
}
