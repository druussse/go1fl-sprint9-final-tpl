package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-10, -1, 0, 1, 8, 1000} {
		t.Run(fmt.Sprintf("size_%d", size), func(t *testing.T) {
			data := generateRandomElements(size)
			wantLength := size
			if wantLength < 0 {
				wantLength = 0
			}
			if len(data) != wantLength {
				t.Fatalf("generateRandomElements(%d): len = %d, want %d", size, len(data), wantLength)
			}
			for i, value := range data {
				if value <= 0 {
					t.Fatalf("data[%d] = %d: expected a positive integer", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"one element", []int{42}, 42},
		{"all equal", []int{7, 7, 7}, 7},
		{"maximum first", []int{99, 4, 10}, 99},
		{"maximum middle", []int{3, 105, 8}, 105},
		{"maximum last", []int{1, 2, 5}, 5},
		{"negative values", []int{-9, -3, -7}, -3},
		{"mixed values", []int{-1, 0, 15, 3}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"one element", []int{123}, 123},
		{"two elements", []int{11, 5}, 11},
		{"fewer than chunks", []int{4, 1, 9, 2, 3}, 9},
		{"exactly eight", []int{1, 4, 3, 9, 8, 5, 2, 7}, 9},
		{"remainder one", []int{1, 2, 3, 4, 5, 6, 7, 8, 100}, 100},
		{"remainder seven", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 100}, 100},
		{"maximum near start", []int{200, 1, 2, 3, 4, 5, 6, 7, 8}, 200},
		{"all equal", []int{6, 6, 6, 6, 6, 6, 6, 6, 6}, 6},
		{"negative short", []int{-9, -3, -7}, -3},
		{"negative long", []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	generator := rand.New(rand.NewSource(42))
	for _, length := range []int{1, 2, 7, 8, 9, 15, 16, 17, 31, 32, 33, 1001, 10_000} {
		data := make([]int, length)
		for i := range data {
			data[i] = generator.Intn(2001) - 1000
		}
		if got, want := maxChunks(data), maximum(data); got != want {
			t.Fatalf("length %d: maxChunks = %d, maximum = %d", length, got, want)
		}
	}
}
