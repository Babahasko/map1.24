package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a large map
	m := make(map[int]int, 1_000_000)
	for i := 0; i < 1_000_000; i++ {
		m[i] = i
	}

	// Measure lookup performance
	start := time.Now()
	for i := 0; i < 10_000_000; i++ {
		_ = m[i%1_000_000]
	}
	fmt.Printf("Lookup time: %v\n", time.Since(start))

	// Measure insertion performance
	start = time.Now()
	for i := 1_000_000; i < 2_000_000; i++ {
		m[i] = i
	}
	fmt.Printf("Insertion time: %v\n", time.Since(start))

	// Measure deletion performance
	start = time.Now()
	for i := 0; i < 1_000_000; i++ {
		delete(m, i)
	}
	fmt.Printf("Deletion time: %v\n", time.Since(start))
}
