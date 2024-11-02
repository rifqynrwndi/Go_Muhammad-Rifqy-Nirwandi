package main

import (
	"testing"
)

func TestRectangleArea(t *testing.T) {
	tests := []struct {
		length, width int
		expected      string
	}{
		{4, 5, "even rectangle"},
		{3, 5, "odd rectangle"},
		{2, 2, "even rectangle"},
		{1, 1, "odd rectangle"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := RectangleArea(tt.length, tt.width)
			if got != tt.expected {
				t.Errorf("RectangleArea(%d, %d) = %v; want %v", tt.length, tt.width, got, tt.expected)
			}
		})
	}
}

func TestRectanglePerimeter(t *testing.T) {
	tests := []struct {
		length, width int
		expected      bool
	}{
		{4, 5, false},
		{10, 10, true},
		{1, 1, false},
		{3, 2, true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := RectanglePerimeter(tt.length, tt.width)
			if got != tt.expected {
				t.Errorf("RectanglePerimeter(%d, %d) = %v; want %v", tt.length, tt.width, got, tt.expected)
			}
		})
	}
}

func TestSquareArea(t *testing.T) {
	tests := []struct {
		s       int
		expected string
	}{
		{4, "even square"},
		{3, "odd square"},
		{2, "even square"},
		{1, "odd square"},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := SquareArea(tt.s)
			if got != tt.expected {
				t.Errorf("SquareArea(%d) = %v; want %v", tt.s, got, tt.expected)
			}
		})
	}
}

func TestSquarePerimeter(t *testing.T) {
	tests := []struct {
		s       int
		expected bool
	}{
		{10, true},
		{8, false},
		{9, false},
		{11, true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := SquarePerimeter(tt.s)
			if got != tt.expected {
				t.Errorf("SquarePerimeter(%d) = %v; want %v", tt.s, got, tt.expected)
			}
		})
	}
}
