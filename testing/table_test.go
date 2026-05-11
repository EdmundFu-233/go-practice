package main

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b     float64
		expected float64
		hasError bool
	}{
		{10, 2, 5, false},
		{5, 0, 0, true},
		{-6, 3, -2, false},
	}
	for _, tt := range tests {
		got, err := divide(tt.a, tt.b)
		if tt.hasError && err == nil {
			t.Errorf("divide(%f,%f) expected error", tt.a, tt.b)
		}
		if !tt.hasError && got != tt.expected {
			t.Errorf("divide(%f,%f) = %f; want %f", tt.a, tt.b, got, tt.expected)
		}
	}
}
