package unittest

import (
	"testing"
	"time"
)

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(b, 1)
}
func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}
func BenchmarkFib3(b *testing.B) {
	benchmarkFib(b, 3)
}
func BenchmarkFib10(b *testing.B) {
	benchmarkFib(b, 10)
}
func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}
func BenchmarkFib40(b *testing.B) {
	benchmarkFib(b, 40)
}

func BenchmarkSplit(b *testing.B) {
	time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
	b.ResetTimer()              // 重置计时器
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

func BenchmarkSplitParallel(b *testing.B) {
	// b.SetParallelism(1) // 设置使用的CPU数
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Split("沙河有沙又有河", "沙")
		}
	})
}
