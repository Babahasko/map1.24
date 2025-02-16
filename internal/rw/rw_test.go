package rw

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkReadFromEmptyMap(b *testing.B) {
	d := &Data{M: make(map[string]any)}
	for i := 0; i < b.N; i++ {
		_, _ = d.ReadFromMap("value")
	}
}

func BenchmarkReadFromMap(b *testing.B) {
	d := &Data{M: make(map[string]any, 10_000_000)}

	// Populate the map with 10,000,000 elements
	for i := 0; i < 10_000_000; i++ {
		d.M[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", rand.Intn(10_000_000))
		_, _ = d.ReadFromMap(key)
	}
}

func BenchmarkWriteToUnallocatedMap(b *testing.B) {
	// Unallocated map
	d := &Data{M: make(map[string]any, 0)}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key%d", i)
		d.M[key] = fmt.Sprintf("value%d", i)
	}
}
