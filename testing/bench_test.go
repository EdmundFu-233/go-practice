package main

import "testing"

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(20)
	}
}

func BenchmarkFibOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibOptimized(20)
	}
}
