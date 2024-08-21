package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// Definisikan tabel test dengan berbagai kasus uji
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Positive numbers", 1, 2, 3},
		{"Negative numbers", -1, -1, -2},
		{"Mix of positive and negative", 1, -1, 0},
		{"Zero", 0, 0, 0},
	}

	// Iterasi melalui setiap kasus uji dalam tabel test
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
