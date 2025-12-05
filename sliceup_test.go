package sliceup

import (
	"testing"
)

func BenchmarkUpdateWithAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UpdateWithAppend()
	}
}

func BenchmarkUpdateWithIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UpdateWithIndex()
	}
}

func BenchmarkUpdateArrayWithIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UpdateArrayWithIndex()
	}
}
