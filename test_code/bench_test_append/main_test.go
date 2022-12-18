package main

import "testing"

func BenchmarkAppendZeroSlice(b *testing.B) {
	maxNum := 10000
	for i := 0; i < b.N; i++ {
		var zeroSlice []int
		for j := 0; j < maxNum; j++ {
			zeroSlice = append(zeroSlice, j)
		}
	}
}

func BenchmarkAppendCapSlice(b *testing.B) {
	maxNum := 10000
	for i := 0; i < b.N; i++ {
		capSlice := make([]int, 0, maxNum)
		for j := 0; j < maxNum; j++ {
			capSlice = append(capSlice, j)
		}
	}
}
