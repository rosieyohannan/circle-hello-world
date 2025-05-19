package main

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "TestAdd_PositiveNumbers",
			a:        1,
			b:        2,
			expected: 3,
		},
		{
			name:     "TestAdd_NegativeNumbers",
			a:        -1,
			b:        -2,
			expected: -3,
		},
		{
			name:     "TestAdd_MixedNumbers",
			a:        -1,
			b:        2,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simulate some work
			time.Sleep(100 * time.Millisecond)
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestMultiplication(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "TestMultiplication_PositiveNumbers",
			a:        2,
			b:        3,
			expected: 6,
		},
		{
			name:     "TestMultiplication_WithZero",
			a:        0,
			b:        5,
			expected: 0,
		},
		{
			name:     "TestMultiplication_NegativeNumbers",
			a:        -2,
			b:        3,
			expected: -6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simulate some work
			time.Sleep(150 * time.Millisecond)
			result := tt.a * tt.b
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestDivision(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
		wantErr  bool
	}{
		{
			name:     "TestDivision_ValidDivision",
			a:        6,
			b:        2,
			expected: 3,
			wantErr:  false,
		},
		{
			name:     "TestDivision_DivideByZero",
			a:        5,
			b:        0,
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "TestDivision_NegativeNumbers",
			a:        -6,
			b:        2,
			expected: -3,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simulate some work
			time.Sleep(200 * time.Millisecond)
			if tt.b == 0 {
				if !tt.wantErr {
					t.Error("expected error for division by zero")
				}
				return
			}
			result := tt.a / tt.b
			if result != tt.expected {
				t.Errorf("got %d, want %d", result, tt.expected)
			}
		})
	}
}