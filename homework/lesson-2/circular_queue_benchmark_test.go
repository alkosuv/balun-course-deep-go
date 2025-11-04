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

// Benchmark для последовательных операций Push/Pop
func BenchmarkCircularQueueWithoutLen_PushPop_10_000(b *testing.B) {
	size := 10_000
	queue := NewCircularQueueWithoutLen[int](size)

	b.ResetTimer()
	for b.Loop() {
		for j := 0; j <= size; j++ {
			queue.Push(j)
		}

		for j := 0; j <= size; j++ {
			queue.Pop()
		}
	}
}

// Benchmark для последовательных операций Push/Pop
func BenchmarkCircularQueue_PushPop_10_000(b *testing.B) {
	size := 10_000
	queue := NewCircularQueue[int](size)

	b.ResetTimer()
	for b.Loop() {
		for j := 0; j <= size; j++ {
			queue.Push(j)
		}

		for j := 0; j <= size; j++ {
			queue.Pop()
		}
	}
}

// Benchmark для последовательных операций Push/Pop
func BenchmarkCircularQueueWithoutLen_PushPop_100_000(b *testing.B) {
	size := 100_000
	queue := NewCircularQueueWithoutLen[int](size)

	b.ResetTimer()
	for b.Loop() {
		for j := 0; j <= size; j++ {
			queue.Push(j)
		}

		for j := 0; j <= size; j++ {
			queue.Pop()
		}
	}
}

// Benchmark для последовательных операций Push/Pop
func BenchmarkCircularQueue_PushPop_100_000(b *testing.B) {
	size := 100_000
	queue := NewCircularQueue[int](size)

	b.ResetTimer()
	for b.Loop() {
		for j := 0; j <= size; j++ {
			queue.Push(j)
		}

		for j := 0; j <= size; j++ {
			queue.Pop()
		}
	}
}

// Benchmark для последовательных операций Push/Pop
func BenchmarkCircularQueueWithoutLen_PushPop_1_000_000(b *testing.B) {
	size := 1_000_000
	queue := NewCircularQueueWithoutLen[int](size)

	b.ResetTimer()
	for b.Loop() {
		for j := 0; j <= size; j++ {
			queue.Push(j)
		}

		for j := 0; j <= size; j++ {
			queue.Pop()
		}
	}
}

// Benchmark для последовательных операций Push/Pop
func BenchmarkCircularQueue_PushPop_1_000_000(b *testing.B) {
	size := 1_000_000
	queue := NewCircularQueue[int](size)

	b.ResetTimer()
	for b.Loop() {
		for j := 0; j <= size; j++ {
			queue.Push(j)
		}

		for j := 0; j <= size; j++ {
			queue.Pop()
		}
	}
}

// Benchmark для последовательных операций Push/Pop
func BenchmarkCircularQueueWithoutLen_PushPop_1_000_000_000(b *testing.B) {
	size := 1_000_000_000
	queue := NewCircularQueueWithoutLen[int](size)

	b.ResetTimer()
	for b.Loop() {

		for j := 0; j <= size; j++ {
			queue.Push(j)
		}

		for j := 0; j <= size; j++ {
			queue.Pop()
		}
	}
}

// Benchmark для последовательных операций Push/Pop
func BenchmarkCircularQueue_PushPop_1_000_000_000(b *testing.B) {
	size := 1_000_000_000
	queue := NewCircularQueue[int](size)

	b.ResetTimer()
	for b.Loop() {
		for j := 0; j <= size; j++ {
			queue.Push(j)
		}

		for j := 0; j <= size; j++ {
			queue.Pop()
		}
	}
}
