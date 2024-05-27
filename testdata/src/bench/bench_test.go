package bench

import (
	"fmt"
	"strconv"
	"testing"
)

func TestOne(t *testing.T) {
	t.Logf("test log")
}

func BenchmarkOne(b *testing.B) {
	b.Logf("benchmark log (%d)", b.N)
	for i := 0; i < b.N; i++ {
		test(i)
	}
}

func BenchmarkTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		n, _ := strconv.Atoi(strconv.Itoa(i))
		test(n)
	}
}

func test(x int) int {
	return x + 1
}

func BenchmarkTree(b *testing.B) {
	sizes := []int{1, 2, 4, 8, 16}
	for _, size := range sizes {
		b.Run(
			fmt.Sprintf("size: %d", size),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					test(i + size)
				}
			})
	}
}
