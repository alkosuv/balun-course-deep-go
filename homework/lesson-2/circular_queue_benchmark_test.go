package main

import "testing"

// Benchmark для последовательных операций Push/Pop
func BenchmarkCircularQueueWithoutLen_PushPop(b *testing.B) {
	queue := NewCircularQueueWithoutLen[int](b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.Push(i)
		queue.Pop()
	}
}

// Benchmark для последовательных операций Push/Pop
func BenchmarkCircularQueue_PushPop(b *testing.B) {
	queue := NewCircularQueue[int](b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		queue.Push(i)
		queue.Pop()
	}
}
