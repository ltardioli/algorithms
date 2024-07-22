package utils

import "testing"

func BenchmarkEnqueue(b *testing.B) {
	b.Run("1000 elements", func(b *testing.B) {
		var q Queue[int]
		for i := range 1000 {
			q.Enqueue(i)
		}
	})

	b.Run("10000 elements", func(b *testing.B) {
		var q Queue[int]
		for i := range 10000 {
			q.Enqueue(i)
		}
	})

	b.Run("100000 elements", func(b *testing.B) {
		var q Queue[int]
		for i := range 100000 {
			q.Enqueue(i)
		}
	})

	b.Run("1000000 elements", func(b *testing.B) {
		var q Queue[int]
		for i := range 1000000 {
			q.Enqueue(i)
		}
	})

	b.Run("100000000 elements", func(b *testing.B) {
		var q Queue[int]
		for i := range 100000000 {
			q.Enqueue(i)
		}
	})

	b.Run("1000000000 elements", func(b *testing.B) {
		var q Queue[int]
		for i := range 1000000000 {
			q.Enqueue(i)
		}
	})

}

func BenchmarkDequeue(b *testing.B) {
	b.Run("1000 elements", func(b *testing.B) {
		var q Queue[int]
		for i := range 1000 {
			q.Enqueue(i)
		}

		for !q.IsEmpty() {
			q.Dequeue()
		}
	})

	b.Run("10000 elements", func(b *testing.B) {
		var q Queue[int]
		for i := range 10000 {
			q.Enqueue(i)
		}

		for !q.IsEmpty() {
			q.Dequeue()
		}
	})

	b.Run("100000 elements", func(b *testing.B) {
		var q Queue[int]
		for i := range 100000 {
			q.Enqueue(i)
		}

		for !q.IsEmpty() {
			q.Dequeue()
		}
	})

	b.Run("1000000 elements", func(b *testing.B) {
		var q Queue[int]
		for i := range 1000000 {
			q.Enqueue(i)
		}

		for !q.IsEmpty() {
			q.Dequeue()
		}
	})
}
