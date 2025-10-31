package main

type CircularQueue[T int8 | int16 | int32 | int64 | int] struct {
	queue []T
	head  int
	tail  int
	len   int
	cap   int
}

// NewCircularQueue создать очередь с определенным размером буффера
func NewCircularQueue[T int8 | int16 | int32 | int64 | int](size int) *CircularQueue[T] {
	return &CircularQueue[T]{
		queue: make([]T, size),
		cap:   size,
		len:   0,
		head:  0,
		tail:  0,
	}
}

// Dequeue используется для извлечения и одновременного удаления значения из начала очереди
func (q *CircularQueue[T]) Dequeue() (T, bool) {
	return q.Front(), q.Pop()
}

// Push добавить значение в конец очереди (false, если очередь заполнена)
func (q *CircularQueue[T]) Push(value T) bool {
	if q.Full() {
		return false
	}

	q.queue[q.tail] = value
	q.tail = (q.tail + 1) % q.cap
	q.len++
	return true
}

// Pop удалить значение из начала очереди (false, если очередь пустая)
func (q *CircularQueue[T]) Pop() bool {
	if q.Empty() {
		return false
	}

	q.head = (q.head + 1) % q.cap
	q.len--
	return true
}

// Front получить значение из начала очереди (-1, если очередь пустая)
func (q *CircularQueue[T]) Front() T {
	if q.Empty() {
		return -1
	}
	return q.queue[q.head]
}

// Back получить значение из конца очереди (-1, если очередь пустая)
func (q *CircularQueue[T]) Back() T {
	if q.Empty() {
		return -1
	}

	// Вычисляем позицию последнего элемента
	tail := (q.tail - 1 + q.cap) % q.cap
	return q.queue[tail]
}

// Empty проверить пустая ли очередь
func (q *CircularQueue[T]) Empty() bool {
	return q.len == 0
}

// Full проверить заполнена ли очередь
func (q *CircularQueue[T]) Full() bool {
	return q.len == q.cap
}

// getQueue получить очередь
func (q *CircularQueue[T]) getQueue() []T {
	return q.queue
}
